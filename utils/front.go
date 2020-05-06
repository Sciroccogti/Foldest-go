package utils

import (
	"foldest-go/lorca"
	"sync"
)

type Front struct {
	sync.Mutex
	isStart bool
	Ui      lorca.UI
}

func mainThread(pf *Front, exitsignal chan<- bool) {
	pf.Ui.Eval(`console.log("Starting...");`)
	Plog.Print("Starting...\n")
	conf := ReadConf()

	rules := ReadRules()
	if rules == nil {
		Plog.Print("Skipping classify...\n")
	} else {
		DoClassify(rules, conf.Targetdir, conf.Verbose)
	}

	if conf.Tmpbin.Enable {
		Plog.Print("Performing tmpbin...\n")
		Manage(conf)
	} else {
		Plog.Print("tmpbin is disabled, skipping...\n")
	}

	Plog.Print("Exiting...\n")
	pf.Ui.Eval(`console.log("Exiting...");`)
	exitsignal <- true
}

func (pf *Front) Start() {
	pf.Lock()
	defer pf.Unlock()
	exitsignal := make(chan bool)
	pf.isStart = true
	go mainThread(pf, exitsignal)
	<-exitsignal
	pf.isStart = false
}

func (pf *Front) Stop() {
	pf.Lock()
	defer pf.Unlock()
	pf.isStart = false
}

func (pf *Front) Status() (status bool) {
	// pf.Lock()
	// defer pf.Unlock()
	return pf.isStart
}
