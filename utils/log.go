package utils

import (
	"fmt"
	"log"
	"os"
	"time"
)

const logPath = "/usr/local/cni/log"

func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func Write(path, format string, v ...string) {
	dateTime := time.Now().Format(
		"2006-01-02",
	)
	var basePath = ""
	if path != "" {
		basePath = path
	} else {
		basePath = logPath
	}
	if !Exists(basePath) {
		os.MkdirAll(basePath, os.ModePerm)
	}
	writer, error := os.OpenFile(basePath+"/"+dateTime+".log",
		os.O_APPEND|os.O_CREATE|os.O_RDWR, os.ModePerm)
	if error != nil {
		fmt.Printf(error.Error())
		os.Exit(1)
		return
	}
	defer writer.Close()

	Logger := log.New(writer, "", log.Ltime|log.Lshortfile|log.LstdFlags)
	Logger.Println(fmt.Sprintf(format, v))
}
