package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"
)

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

// CheckDir : Check if path is a folder
func CheckDir(path string) (isDir bool) {
	info, err := os.Stat(path)
	if err != nil || !info.IsDir() {
		fmt.Printf("Error while scanning %c[0;34m%s%c[0m :", 0x1B, path, 0x1B)
		fmt.Printf("\t%c[0;31m%s%c[0m\n", 0x1B, err, 0x1B)
		return false
	}
	return true
}

//OpenDir : Open a folder
func OpenDir(path string) (dir []os.FileInfo) {
	fmt.Printf("Scanning %c[0;34m%s%c[0m ...\n", 0x1B, path, 0x1B)
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Printf("Error while scanning %c[0;34m%s%c[0m :", 0x1B, path, 0x1B)
		fmt.Printf("\t%c[0;31m%s%c[0m\n", 0x1B, err, 0x1B)
		return nil
	}
	return dir
}
