package main

import (
	"fmt"
	"foldest-go/utils"
	"time"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func main() {
	err := ui.Main(func() {
		startBtn := ui.NewButton("Start")
		startBtn.OnClicked(func(*ui.Button) {
			startTime := time.Now().UnixNano()

			conf := utils.ReadConf()

			rules := utils.ReadRules()
			if rules == nil {
				fmt.Println("Skipping classify...")
			} else {
				utils.DoClassify(rules, conf.Targetdir, conf.Verbose)
			}

			if conf.Tmpbin.Enable {
				fmt.Println("Performing tmpbin...")
				utils.Manage(conf)
			} else {
				fmt.Println("tmpbin is disabled, skipping...")
			}

			endTime := time.Now().UnixNano()
			seconds := float64((endTime - startTime) / 1e9)
			fmt.Printf("Finished in %.2f sec.\n", seconds)
		})

		outputLabel := ui.NewLabel("")

		canvas := ui.NewVerticalBox()
		canvas.Append(startBtn, false)
		canvas.Append(outputLabel, true)
		window := ui.NewWindow("Foldest", 300, 300, false)
		window.SetChild(canvas)
		window.SetMargined(true)
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()
	})
	if err != nil {
		panic(err)
	}
}
