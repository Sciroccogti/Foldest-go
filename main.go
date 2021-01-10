package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"foldest-go/utils"
	"os"
	"time"

	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
	bootstrap "github.com/asticode/go-astilectron-bootstrap"
)

// Constants
const version = `3.0`

// Vars injected via ldflags by bundler
var (
	AppName            string
	BuiltAt            string
	VersionAstilectron string
	VersionElectron    string
)

// Application Vars
var (
	fs    = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	debug = fs.Bool("d", false, "enables the debug mode")
)

func main() {
	os.MkdirAll("./logs", 0766)

	var aboutPayload  map[string]string
	aboutPayload = make(map[string]string)
	aboutPayload["AppName"] =      AppName
	aboutPayload["Version"] =      version
	aboutPayload["BuiltTime"] =    BuiltAt[:10]
	aboutPayload["Electron"] =     VersionElectron
	aboutPayload["Astilectron"] =  VersionAstilectron
	aboutPayload["Github"] =       "https://github.com/Sciroccogti/Foldest-go"

	// Parse flags
	fs.Parse(os.Args[1:])

	// Run bootstrap
	utils.L.Printf("Running app built at %s\n", BuiltAt)
	if err := bootstrap.Run(bootstrap.Options{
		Asset:    Asset,
		AssetDir: AssetDir,
		AstilectronOptions: astilectron.Options{
			AppName:            AppName,
			AppIconDarwinPath:  "resources/icon.icns",
			AppIconDefaultPath: "resources/icon.png",
			SingleInstance:     true,
			VersionAstilectron: VersionAstilectron,
			VersionElectron:    VersionElectron,
		},
		Debug:  *debug,
		Logger: utils.L,
		MenuOptions: []*astilectron.MenuItemOptions{
			{
				Label: astikit.StrPtr("File"),
				SubMenu: []*astilectron.MenuItemOptions{
					{Role: astilectron.MenuItemRoleClose},
				},
			},
			{
				Label: astikit.StrPtr("Help"),
				SubMenu: []*astilectron.MenuItemOptions{
					{Role: astilectron.MenuItemRoleToggleDevTools},
					{
						Label: astikit.StrPtr("About"),
						OnClick: func(e astilectron.Event) (deleteListener bool) {
							if err := bootstrap.SendMessage(utils.W, "about", aboutPayload, func(m *bootstrap.MessageIn) {
								// Unmarshal payload
								var s string
								if err := json.Unmarshal(m.Payload, &s); err != nil {
									utils.L.Println(fmt.Errorf("unmarshaling payload failed: %w", err))
									return
								}
								utils.L.Printf("About modal has been displayed and payload is %s!\n", s)
							}); err != nil {
								utils.L.Println(fmt.Errorf("sending about event failed: %w", err))
							}
							return
						},
					},
				},
			},
		},
		OnWait: func(_ *astilectron.Astilectron, ws []*astilectron.Window, _ *astilectron.Menu, _ *astilectron.Tray, _ *astilectron.Menu) error {
			utils.W = ws[0]
			go func() {
				time.Sleep(5 * time.Second)

				if err := bootstrap.SendMessage(utils.W, "check.out.menu", "Don't forget to check out the menu!"); err != nil {
					utils.L.Println(fmt.Errorf("sending check.out.menu event failed: %w", err))
				}
			}()
			return nil
		},
		RestoreAssets: RestoreAssets,
		Windows: []*bootstrap.Window{{
			Homepage:       "index.html",
			MessageHandler: handleMessages,
			Options: &astilectron.WindowOptions{
				BackgroundColor: astikit.StrPtr("#333"),
				Center:          astikit.BoolPtr(true),
				Height:          astikit.IntPtr(700),
				Width:           astikit.IntPtr(700),
			},
		}},
	}); err != nil {
		utils.L.Fatal(fmt.Errorf("running bootstrap failed: %w", err))
	}
}
