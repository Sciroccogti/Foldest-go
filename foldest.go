//go:generate go run -tags generate gen.go

package main

import (
	"fmt"
	"foldest-go/utils"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"sync"

	"foldest-go/lorca"
)

type start struct {
	sync.Mutex
	isStart bool
	ui      lorca.UI
}

func (s *start) Start() {
	s.ui.Eval(`console.log("OHHHHHHHHHH!");`)
	s.Lock()
	defer s.Unlock()
	exitsignal := make(chan bool)
	s.isStart = true
	go mainThread(s, exitsignal)
	<-exitsignal
	s.ui.Eval(`console.log("!!!!!!!!!!!!!!!!!!!!!!!!!!!");`)
	s.isStart = false
}

func (s *start) Stop() {
	s.Lock()
	defer s.Unlock()
	s.isStart = false
}

func (s *start) Status() (status bool) {
	s.Lock()
	defer s.Unlock()
	return s.isStart
}

func (s *start) Output() (output string) {
	s.Lock()
	defer s.Unlock()
	return "haha!\n"
}

func main() {
	args := []string{}
	if runtime.GOOS == "linux" {
		args = append(args, "--class=Lorca")
	}
	s := &start{}
	var err error
	s.ui, err = lorca.New("", "", 480, 320, args...)
	if err != nil {
		log.Fatal(err)
	}
	defer s.ui.Close()

	// A simple way to know when UI is ready (uses body.onload event in JS)
	s.ui.Bind("start", func() {
		log.Println("UI is ready")
	})

	// Create and bind Go object to the UI
	s.ui.Bind("mainStart", s.Start)
	// ui.Bind("mainStop", s.Stop)
	s.ui.Bind("mainStatus", s.Status)
	s.ui.Bind("getoutput", s.Output)

	// Load HTML.
	// You may also use `data:text/html,<base64>` approach to load initial HTML,
	// e.g: ui.Load("data:text/html," + url.PathEscape(html))

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	go http.Serve(ln, http.FileServer(FS))
	s.ui.Load(fmt.Sprintf("http://%s", ln.Addr()))

	// You may use console.log to debug your JS code, it will be printed via
	// log.Println(). Also exceptions are printed in a similar manner.
	s.ui.Eval(`
		console.log("Hello, world!");
		console.log('Multiple values:', [1, false, {"x":5}]);
	`)

	// Wait until the interrupt signal arrives or browser window is closed
	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)
	select {
	case <-sigc:
	case <-s.ui.Done():
	}

	log.Println("exiting...")
}

func mainThread(s *start, exitsignal chan<- bool) {
	s.ui.Eval(`console.log("Starting...");`)
	fmt.Println("Starting...")
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

	fmt.Println("Exiting...")
	s.ui.Eval(`console.log("Exiting...");`)
	exitsignal <- true
}
