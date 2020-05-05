package main

import (
	"fmt"
	"foldest-go/utils"
)

func main() {
	conf := utils.ReadConf()
	fmt.Println("Press enter to start...")
	fmt.Scanln()

	if conf.Tmpbin.Enable {
		utils.Manage(conf)
	} else {
		fmt.Println("tmpbin is disabled, skipping...")
	}

	fmt.Println("Press enter to exit...")
	fmt.Scanln()
}
