package log

import (
	"PWL-Client/conf"
	"io"
	"log"
	"os"
	"strings"
)

var (
	Info    *log.Logger
	Message *log.Logger
)

func init() {
	prefix := conf.GetVal("profile", "username")
	//log.SetPrefix(prefix)
	//log.SetFlags(log.Ltime | log.Lshortfile)
	msgLog, err1 := os.OpenFile("msg.log", os.O_WRONLY|os.O_CREATE, 0666)
	if err1 != nil {
		log.Println("打开msg文件错误")
	}
	infoLog, err2 := os.OpenFile("info.log", os.O_WRONLY|os.O_CREATE, 0666)
	if err2 != nil {
		log.Println("打开info文件错误")
	}

	Info = log.New(io.MultiWriter(infoLog, os.Stdout),strings.Join([]string{prefix,":"},""), log.Lshortfile)
	Message = log.New(io.MultiWriter(msgLog, os.Stdout), strings.Join([]string{prefix,":"},""), log.Lshortfile)
}
