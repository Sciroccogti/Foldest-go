package main

import (
	"fmt"
	"foldest-go/utils"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"

	yaml "gopkg.in/yaml.v2"
)

func main() {
	// Read conf.yml
	fmt.Println("Reading conf.yml ...")
	conf := new(utils.Yaml)
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
	var path string
	path = conf.Targetdir
	isChanged := false

	for path == "" {
		fmt.Println("Please input path of the target folder:")
		fmt.Scanln(&path)
		if !strings.HasSuffix(path, "/") {
			path = path + "/"
		}
		info, err := os.Stat(path)
		if err != nil || !info.IsDir() {
			fmt.Printf("Error while scanning %c[0;34m%s%c[0m :", 0x1B, path, 0x1B)
			fmt.Printf("\t%c[0;31m%s%c[0m\n", 0x1B, err, 0x1B)
			path = ""
			continue
		}
		conf.Targetdir = path
		isChanged = true
	}

	// Set default values
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

	// Scan target dir
	fmt.Printf("Scanning %c[0;34m%s%c[0m ...\n", 0x1B, path, 0x1B)
	if !strings.HasSuffix(path, "/") {
		path = path + "/"
	}
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Printf("Error while scanning %c[0;34m%s%c[0m :", 0x1B, path, 0x1B)
		fmt.Printf("\t%c[0;31m%s%c[0m\n", 0x1B, err, 0x1B)
		return
	}

	fmt.Println("Press any key to start...")
	fmt.Scanln()

	// Make dir 'tmpbin/'

	_, err = os.Stat(path + conf.Tmpbin.Name)
	if err != nil {
		fmt.Printf("Making folder %c[0;34m%s%c[0m ...\n", 0x1B, conf.Tmpbin.Name, 0x1B)
		err := os.Mkdir(path+conf.Tmpbin.Name, 0777)
		if err != nil {
			fmt.Printf("Error while making folder %c[0;34m%s%c[0m ...\n", 0x1B, conf.Tmpbin.Name, 0x1B)
			fmt.Printf("\t%c[0;31m%s%c[0m\n", 0x1B, err, 0x1B)
		}
	}

	// Operations on files
	if conf.Tmpbin.Enable {
		for count, file := range dir {
			if count > 10 {
				break
			}
			modTime, strerr := GetFileModTime(path + file.Name())
			if strerr == "" {
				// jump tmpbin
				if file.Name() == conf.Tmpbin.Name {
					continue
				}

				if conf.Verbose {
					fmt.Printf("%c[0;34m%s%c[0m %c[0;32m%s%c[0m\n", 0x1B, file.Name(), 0x1B, 0x1B, modTime, 0x1B)
				}

				// If file reaches thresh
				if time.Now().Unix()-modTime.Unix() >= int64(conf.Tmpbin.Thresh*86400) {
					if conf.Verbose {
						fmt.Printf("Moving %c[0;34m%s%c[0m\n", 0x1B, file.Name(), 0x1B)
					}
					src := path + file.Name()
					des := path + conf.Tmpbin.Name + file.Name()

					// Check if file already existed in tmpbin
					if _, err := os.Stat(des); !os.IsNotExist(err) {
						fmt.Printf("Error while moving %c[0;34m%s%c[0m :", 0x1B, file.Name(), 0x1B)
						fmt.Printf("\t%c[0;31m%s already existed in %s%c[0m\n", 0x1B, file.Name(), conf.Tmpbin.Name, 0x1B)
					} else {
						if !file.IsDir() { // file, not folder
							_, err := CopyFile(src, des)
							if err != nil {
								fmt.Printf("Error while moving %c[0;34m%s%c[0m :", 0x1B, file.Name(), 0x1B)
								fmt.Printf("\t%c[0;31m%s%c[0m\n", 0x1B, err, 0x1B)
							}
							err = os.Remove(src)
							if err != nil {
								fmt.Printf("Error while moving %c[0;34m%s%c[0m :", 0x1B, file.Name(), 0x1B)
								fmt.Printf("\t%c[0;31m%s%c[0m\n", 0x1B, err, 0x1B)
							}
						} else { // folder
							err := os.Rename(src, des)
							if err != nil {
								fmt.Printf("Error while moving %c[0;34m%s%c[0m :", 0x1B, file.Name(), 0x1B)
								fmt.Printf("\t%c[0;31m%s%c[0m\n", 0x1B, err, 0x1B)
							}
						}
					}
				}

			} else {
				fmt.Printf("Error while scanning %c[0;34m%s%c[0m :", 0x1B, file.Name(), 0x1B)
				fmt.Printf("\t%c[0;31m%s%c[0m\n", 0x1B, err, 0x1B)
			}
		}
	}

	// Save conf.yml
	if isChanged {
		fmt.Println("Saving conf.yml ...")
		yamlChanged, err := yaml.Marshal(conf)
		if err != nil {
			fmt.Printf("Error while saving conf.yml :\n")
			fmt.Printf("\t%c[0;31m%s%c[0m\n", 0x1B, err, 0x1B)
		}
		err = ioutil.WriteFile("conf.yml", yamlChanged, 0644)
	}

	fmt.Println("Press any key to exit...")
	fmt.Scanln()

}

// GetFileModTime ：获取文件修改时间 返回时间
func GetFileModTime(path string) (t time.Time, strerr string) {
	f, err := os.Open(path)
	if err != nil {
		return time.Now(), "open file error"
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return time.Now(), "stat fileinfo error"
	}

	return fi.ModTime(), ""
}

// CopyFile : via io.Copy
func CopyFile(src, des string) (written int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer srcFile.Close()

	//获取源文件的权限
	fi, _ := srcFile.Stat()
	perm := fi.Mode()

	//desFile, err := os.Create(des)  //无法复制源文件的所有权限
	desFile, err := os.OpenFile(des, os.O_RDWR|os.O_CREATE|os.O_TRUNC, perm) //复制源文件的所有权限
	if err != nil {
		return 0, err
	}
	defer desFile.Close()

	return io.Copy(desFile, srcFile)
}
