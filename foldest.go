package main

import (
	"fmt"
	"foldest-go/uiparser"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func main() {
	fmt.Println("hello")
	err := ui.Main(uiparser.SetupUI)
	if err != nil {
		panic(err)
	}
}
