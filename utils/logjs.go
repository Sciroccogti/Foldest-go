package utils

import (
	"fmt"
	"regexp"
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
	end := regexp.MustCompile(`.*?%c[0;\dm.*?(%c[0m)`)
	s = end.ReplaceAllString(s, "</span>")
	red := regexp.MustCompile(`.*?(%c[0;31m).*?%c[0m`)
	s = red.ReplaceAllString(s, "<span style=\"color: red;\">")
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
