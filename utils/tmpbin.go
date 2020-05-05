package utils

import (
	"fmt"
	"os"
	"time"
)

// Manage folder
func Manage(conf *Conf) {
	// Make dir 'tmpbin/'
	_, err := os.Stat(conf.Targetdir + conf.Tmpbin.Name)
	if err != nil {
		fmt.Printf("Making folder %c[0;34m%s%c[0m ...\n", 0x1B, conf.Tmpbin.Name, 0x1B)
		err := os.Mkdir(conf.Targetdir+conf.Tmpbin.Name, 0777)
		if err != nil {
			fmt.Printf("Error while making folder %c[0;34m%s%c[0m ...\n", 0x1B, conf.Tmpbin.Name, 0x1B)
			fmt.Printf("\t%c[0;31m%s%c[0m\n", 0x1B, err, 0x1B)
		}
	}

	dir := OpenDir(conf.Targetdir)
	if dir == nil {
		return
	}

	// Moving files to tmpbin
	for count, file := range dir {
		if count > 10 {
			break
		}
		modTime, strerr := GetFileModTime(conf.Targetdir + file.Name())
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
				//if conf.Verbose {
				fmt.Printf("Moving %c[0;34m%s%c[0m\n", 0x1B, file.Name(), 0x1B)
				//}
				src := conf.Targetdir + file.Name()
				des := conf.Targetdir + conf.Tmpbin.Name + file.Name()
				MoveAll(file, src, des)
			}

		} else {
			fmt.Printf("Error while scanning %c[0;34m%s%c[0m :", 0x1B, file.Name(), 0x1B)
			fmt.Printf("\t%c[0;31m%s%c[0m\n", 0x1B, err, 0x1B)
		}
	}

	// Delete files in tmpbin/
	dir = OpenDir(conf.Targetdir + conf.Tmpbin.Name)
	if dir == nil {
		return
	}
	for _, file := range dir {
		modTime, strerr := GetFileModTime(conf.Targetdir + conf.Tmpbin.Name + file.Name())
		if strerr == "" {
			if conf.Verbose {
				fmt.Printf("%c[0;34m%s%c[0m %c[0;32m%s%c[0m\n", 0x1B, file.Name(), 0x1B, 0x1B, modTime, 0x1B)
			}

			// If file reaches deleteday
			if time.Now().Unix()-modTime.Unix() >= int64(conf.Tmpbin.Delete*86400) {
				//if conf.Verbose {
				fmt.Printf("Deleting %c[0;34m%s%c[0m\n", 0x1B, file.Name(), 0x1B)
				//}
				src := conf.Targetdir + conf.Tmpbin.Name + file.Name()
				if file.IsDir() {
					err = os.RemoveAll(src)
				} else {
					err = os.Remove(src)
				}

				if err != nil {
					fmt.Printf("Error while deleting %c[0;34m%s%c[0m :", 0x1B, file.Name(), 0x1B)
					fmt.Printf("\t%c[0;31m%s%c[0m\n", 0x1B, err, 0x1B)
				}
			}
		} else {
			fmt.Printf("Error while scanning %c[0;34m%s%c[0m :", 0x1B, file.Name(), 0x1B)
			fmt.Printf("\t%c[0;31m%s%c[0m\n", 0x1B, err, 0x1B)
		}
	}
}
