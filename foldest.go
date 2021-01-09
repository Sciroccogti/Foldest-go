package main

import (
	"foldest-go/utils"
)

// Foldest : main process of foldest
func Foldest() {
	conf := utils.ReadConf()
	// utils.Print("Press enter to start...\n")
	// fmt.Scanln()

	rules := utils.ReadRules()
	if rules == nil {
		utils.Print("Skipping classify...\n")
	} else {
		utils.DoClassify(rules, conf.Targetdir, conf.Verbose)
	}

	if conf.Tmpbin.Enable {
		utils.Print("Performing tmpbin...\n")
		utils.Manage(conf)
	} else {
		utils.Print("tmpbin is disabled, skipping...\n")
	}

	utils.Print("Finished\n")
	return
}
