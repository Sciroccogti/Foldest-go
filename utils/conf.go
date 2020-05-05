package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

// ReadConf : Read conf.yml
func ReadConf() (conf *Yaml) {
	fmt.Println("Reading conf.yml ...")
	conf = new(Yaml)
	if _, err := os.Stat("conf.yml"); os.IsNotExist(err) {
		fmt.Println("conf.yml not found, starting with default value ...")
	} else {
		yamlFile, err := ioutil.ReadFile("conf.yml")
		if err != nil {
			fmt.Printf("Error while reading conf.yml :\n")
			fmt.Printf("\t%c[0;31m%s%c[0m\n", 0x1B, err, 0x1B)
		}
		err = yaml.Unmarshal(yamlFile, conf)
		if err != nil {
			fmt.Printf("Error while reading conf.yml :\n")
			fmt.Printf("\t%c[0;31m%s%c[0m\n", 0x1B, err, 0x1B)
		}
	}

	// Set path
	var isChanged bool
	isChanged = SetPath(&conf.Targetdir)

	// Set default values
	SetDefault(conf)

	if isChanged {
		SaveConf(conf)
	}

	return conf
}

// SetPath : Set target dir
func SetPath(path *string) (isChanged bool) {
	isChanged = false

	for {
		if *path == "" {
			fmt.Println("Please input path of the target folder:")
			fmt.Scanln(path)
			isChanged = true
		}

		if !strings.HasSuffix(*path, "/") {
			*path = *path + "/"
			isChanged = true
		}

		if CheckDir(*path) {
			break
		}
	}

	return isChanged
}

// SetDefault : Set default value of the conf
func SetDefault(conf *Yaml) {
	if conf.Tmpbin.Name == "" {
		conf.Tmpbin.Name = "tmpbin/"
	}
	if !strings.HasSuffix(conf.Tmpbin.Name, "/") {
		conf.Tmpbin.Name = conf.Tmpbin.Name + "/"
	}

	if conf.Tmpbin.Thresh == 0 {
		conf.Tmpbin.Thresh = 30
	}

	if conf.Tmpbin.Delete == 0 {
		conf.Tmpbin.Delete = 30
	}

	if len(conf.Tmpbin.Ignore) == 0 {
		conf.Tmpbin.Ignore = append(conf.Tmpbin.Ignore, ".accelerate")
	}
}

// SaveConf : Save the conf.yml
func SaveConf(conf *Yaml) {
	fmt.Println("Saving conf.yml ...")
	yamlChanged, err := yaml.Marshal(conf)
	if err != nil {
		fmt.Printf("Error while saving conf.yml :\n")
		fmt.Printf("\t%c[0;31m%s%c[0m\n", 0x1B, err, 0x1B)
	}
	err = ioutil.WriteFile("conf.yml", yamlChanged, 0644)
}
