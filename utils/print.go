package utils

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/asticode/go-astilectron"
	bootstrap "github.com/asticode/go-astilectron-bootstrap"
)

// W is the global window var
var W *astilectron.Window

// save log to `./logs/`
var logFile, _ = os.OpenFile("logs/"+time.Now().Format("20060102150405")+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)

// L print log to .log file
var L = log.New(logFile, log.Prefix(), log.Flags())

// Print : Print a string to Electron
func Print(format string, a ...interface{}) {
	s := fmt.Sprintf(format, a...)
	s = strings.ReplaceAll(s, (string)(0x1B)+"[0m", "</span>")
	s = strings.ReplaceAll(s, (string)(0x1B)+"[0;31m", "<span style=\"color: red;\">")
	s = strings.ReplaceAll(s, (string)(0x1B)+"[0;32m", "<span style=\"color: green;\">")
	s = strings.ReplaceAll(s, (string)(0x1B)+"[0;33m", "<span style=\"color: orange;\">")
	s = strings.ReplaceAll(s, (string)(0x1B)+"[0;34m", "<span style=\"color: blue;\">")

	if err := bootstrap.SendMessage(W, "print", s); err != nil {
		L.Println(fmt.Errorf("sending print event failed: %w", err))
	}
}
