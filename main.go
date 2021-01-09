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

// const htmlAbout = `Welcome on <b>Astilectron</b> demo!<br>
// This is using the bootstrap and the bundler.`

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
	htmlAbout := `<b>` + AppName + `</b> Ver ` + version + `<br>
	Built Time: ` + BuiltAt[:10] + `<br>
	Astilectron ` + VersionAstilectron + `<br>
	Electron ` + VersionElectron + `<br>
	<a href="https://github.com/Sciroccogti/Foldest-go">Github</a>
	`

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
		MenuOptions: []*astilectron.MenuItemOptions{{
			Label: astikit.StrPtr("File"),
			SubMenu: []*astilectron.MenuItemOptions{
				{
					Label: astikit.StrPtr("About"),
					OnClick: func(e astilectron.Event) (deleteListener bool) {
						if err := bootstrap.SendMessage(utils.W, "about", htmlAbout, func(m *bootstrap.MessageIn) {
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
				{Role: astilectron.MenuItemRoleClose},
			},
		}},
		OnWait: func(_ *astilectron.Astilectron, ws []*astilectron.Window, _ *astilectron.Menu, _ *astilectron.Tray, _ *astilectron.Menu) error {
			utils.W = ws[0]
			go func() {
				utils.W.OpenDevTools()

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
