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

var W *astilectron.Window

// Create logger
// file := "./" + time.Now().Format("20060102") + ".log"
var logFile, _ = os.OpenFile("./"+time.Now().Format("20060102150405")+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)

// l := log.New(log.Writer(), log.Prefix(), log.Flags())
var L = log.New(logFile, log.Prefix(), log.Flags())

// Print : Print a string to Electron
func Print(format string, a ...interface{}) {
	s := fmt.Sprintf(format, a...)
	s = strings.ReplaceAll(s, (string)(0x1B)+"[0m", "</span>")
	s = strings.ReplaceAll(s, (string)(0x1B)+"[0;31m", "<span style=\"color: red;\">")
	s = strings.ReplaceAll(s, (string)(0x1B)+"[0;32m", "<span style=\"color: green;\">")
	s = strings.ReplaceAll(s, (string)(0x1B)+"[0;33m", "<span style=\"color: yellow;\">")
	s = strings.ReplaceAll(s, (string)(0x1B)+"[0;34m", "<span style=\"color: blue;\">")
	// green := regexp.MustCompile(`.*?(%c\[0;32m)%s%c\[0m`)
	// s = green.ReplaceAllString(s, "<span style=\"color: green;\">")
	// yellow := regexp.MustCompile(`.*?(%c\[0;33m)%s%c\[0m`)
	// s = yellow.ReplaceAllString(s, "<span style=\"color: yellow;\">")
	// blue := regexp.MustCompile(`.*?(%c\[0;34m)%s%c\[0m`)
	// s = blue.ReplaceAllString(s, "<span style=\"color: blue;\">")
	// *Plog.stdout <- s
	if err := bootstrap.SendMessage(W, "print", s); err != nil {

	}
	// Plog.Buf.WriteString(line)
}
