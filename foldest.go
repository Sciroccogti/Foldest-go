//go:generate go run -tags generate gen.go

package main

import (
	"fmt"
	"foldest-go/lorca"
	"foldest-go/utils"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime"
)

func main() {
	args := []string{}
	if runtime.GOOS == "linux" {
		args = append(args, "--class=Lorca")
	}
	pf := &utils.Front{}
	var err error
	pf.Ui, err = lorca.New("", "", 480, 320, args...)
	if err != nil {
		log.Fatal(err)
	}
	defer pf.Ui.Close()

	stdout := make(chan string, 10)
	utils.Plog.Init(&stdout)
	// A simple way to know when UI is ready (uses body.onload event in JS)
	pf.Ui.Bind("start", func() {
		log.Println("UI is ready")
		utils.Plog.Print("UI is ready\n")
	})

	// Create and bind Go object to the UI
	pf.Ui.Bind("mainStart", pf.Start)
	// Ui.Bind("mainStop", pf.Stop)
	pf.Ui.Bind("mainStatus", pf.Status)
	pf.Ui.Bind("getline", utils.Plog.Getline)

	// Load HTML.
	// You may also use `data:text/html,<base64>` approach to load initial HTML,
	// e.g: Ui.Load("data:text/html," + url.PathEscape(html))

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	go http.Serve(ln, http.FileServer(FS))
	pf.Ui.Load(fmt.Sprintf("http://%s", ln.Addr()))

	// You may use console.log to debug your JS code, it will be printed via
	// log.Println(). Also exceptions are printed in a similar manner.
	pf.Ui.Eval(`
				console.log("Hello, world!");
				console.log('Multiple values:', [1, false, {"x":5}]);
			`)

	// Wait until the interrupt signal arrives or browser window is closed
	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)
	select {
	case <-sigc:
	case <-pf.Ui.Done():
	}

	log.Println("exiting...")
	// stdout := make(chan string, 10)
	// utils.Plog.Init(&stdout)
	// go func() {
	// 	for {
	// 		utils.Plog.Print("Scanning %c[0;34m%s%c[0m ...\n", 0x1B, "path", 0x1B)
	// 		utils.Plog.Print("hehe\n")
	// 	}

	// }()
	// for i := 0; i < 100; i++ {
	// 	lines := utils.Plog.Getline()
	// 	fmt.Printf(lines)
	// }
}
