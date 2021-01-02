package utils

import (
	"fmt"
	"time"
)

func Start() {
	startTime := time.Now().UnixNano()

	conf := ReadConf()

	rules := ReadRules()
	if rules == nil {
		fmt.Println("Skipping classify...")
	} else {
		DoClassify(rules, conf.Targetdir, conf.Verbose)
	}

	if conf.Tmpbin.Enable {
		fmt.Println("Performing tmpbin...")
		Manage(conf)
	} else {
		fmt.Println("tmpbin is disabled, skipping...")
	}

	endTime := time.Now().UnixNano()
	seconds := float64((endTime - startTime) / 1e9)
	fmt.Printf("Finished in %.2f sec.\n", seconds)
}
