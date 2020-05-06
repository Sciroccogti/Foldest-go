package utils

import (
	"fmt"
	"strings"
	"sync"
)

// Logjs : a log method for js to read
type Logjs struct {
	// Buf *bytes.Buffer
	buf    []byte
	stdout *chan string
}

// Plog : global pointer to Logjs
var Plog *Logjs

// mu : instance mutex
var mu sync.Mutex

// Init : Init Logjs
func (pl *Logjs) Init(std *chan string) {
	mu.Lock()
	defer mu.Unlock()

	if Plog == nil {
		Plog = &Logjs{}
	}
	Plog.stdout = std
	// Plog.buf []byte = make([]byte, 4096)
}

// Print : Print a string to Logjs
func (pl *Logjs) Print(format string, a ...interface{}) {
	// mu.Lock()
	// defer mu.Unlock()
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
	*Plog.stdout <- s
	// Plog.Buf.WriteString(line)
}

// Getline : Get a string from Logjs
func (pl *Logjs) Getline() (lines string) {
	// mu.Lock()
	// defer mu.Unlock()

	// lines, _ = Plog.Buf.ReadString('\n')
	return <-*Plog.stdout
}
