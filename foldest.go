package main

import (
	"fmt"
	"foldest-go/utils"
)

func main() {
	conf := utils.ReadConf()
	fmt.Println("Press enter to start...")
	fmt.Scanln()

	rules := utils.ReadRules()
	if rules == nil {
		fmt.Println("Skipping classify...")
	} else {
		utils.DoClassify(rules, conf.Targetdir, conf.Verbose)
	}

	if conf.Tmpbin.Enable {
		utils.Manage(conf)
	} else {
		fmt.Println("tmpbin is disabled, skipping...")
	}

	fmt.Println("Press enter to exit...")
	fmt.Scanln()
}
