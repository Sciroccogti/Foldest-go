package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"regexp"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

// ReadRules : Read rules.yml
func ReadRules() (rules *Rules) {
	Plog.Print("Reading rules.yml ...\n")
	rules = new(Rules)

	if _, err := os.Stat("rules.yml"); os.IsNotExist(err) {
		Plog.Print("rules.yml not found, skipping ...\n")
		return nil
	}

	yamlFile, err := ioutil.ReadFile("rules.yml")
	if err != nil {
		fmt.Printf("Error while reading rules.yml :\n")
		fmt.Printf("\t%c[0;31m%s%c[0m\n", 0x1B, err, 0x1B)
	}
	err = yaml.Unmarshal(yamlFile, rules)
	if err != nil {
		fmt.Printf("Error while reading rules.yml :\n")
		fmt.Printf("\t%c[0;31m%s%c[0m\n", 0x1B, err, 0x1B)
	}

	return rules
}

// DoClassify :
func DoClassify(rules *Rules, path string, isVerbose bool) {
	rType := reflect.TypeOf(rules)
	rVal := reflect.ValueOf(rules)
	if rType.Kind() == reflect.Ptr {
		// 传入的rules是指针，需要.Elem()取得指针指向的value
		rType = rType.Elem()
		rVal = rVal.Elem()
	} else {
		panic("rules must be ptr to struct")
	}
	for i := 0; i < rType.NumField(); i++ {
		rule := rVal.Field(i).Interface().(Rule)
		if rule.Enable {
			doRule(&rule, path, isVerbose)
		}
	}
}

// doRule :
func doRule(rule *Rule, path string, isVerbose bool) {
	Plog.Print("Performing rule %c[0;33m%s%c[0m ...\n", 0x1B, rule.Name, 0x1B)
	if !strings.HasSuffix(rule.Name, "/") {
		rule.Name = rule.Name + "/"
	}
	_, err := os.Stat(path + rule.Name)
	if err != nil {
		fmt.Printf("Making folder %c[0;33m%s%c[0m ...\n", 0x1B, rule.Name, 0x1B)
		err := os.Mkdir(path+rule.Name, 0777)
		if err != nil {
			fmt.Printf("Error while making folder %c[0;33m%s%c[0m ...\n", 0x1B, rule.Name, 0x1B)
			fmt.Printf("\t%c[0;31m%s%c[0m\n", 0x1B, err, 0x1B)
		}
	}

	dir := OpenDir(path)
	if dir == nil {
		return
	}

	// Moving files to rule dir
	for _, file := range dir {
		// jump rule dir
		if file.Name() == rule.Name || file.IsDir() {
			continue
		}

		for _, pattern := range rule.Regex {
			re := regexp.MustCompile(pattern)
			match := re.MatchString(file.Name())
			if match {
				modTime, strerr := GetFileModTime(path + file.Name())
				if strerr == "" {
					if isVerbose {
						fmt.Printf("%c[0;34m%s%c[0m %c[0;32m%s%c[0m %d\n", 0x1B, file.Name(), 0x1B, 0x1B, modTime, 0x1B, file.Size())
					}
					// If file reaches deleteday
					if time.Now().Unix()-modTime.Unix() >= int64(rule.Thresh*86400) {
						if (rule.Maxsize <= 0 || file.Size() < (int64)(rule.Maxsize)*1024*1024) && file.Size() > (int64)(rule.Minsize)*1024*1024 {
							if isVerbose {
								fmt.Printf("%c[0;34m%s%c[0m matches %c[0;33m%s%c[0m\n", 0x1B, file.Name(), 0x1B, 0x1B, rule.Name, 0x1B)
							}
							src := path + file.Name()
							des := path + rule.Name + file.Name()
							MoveAll(file, src, des)
						}
					}
				} else {
					fmt.Printf("Error while scanning %c[0;34m%s%c[0m :", 0x1B, file.Name(), 0x1B)
					fmt.Printf("\t%c[0;31m%s%c[0m\n", 0x1B, err, 0x1B)
				}
			}
		}
	}
}
