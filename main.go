package main

import (
	"PWL-Client/conf"
	"PWL-Client/log"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

/*
nameOrEmail：用户名或邮箱地址
userPassword：使用MD5加密后的密码，注意：这里不支持直接传递明文，必须是32位小写MD5加密后的密码
*/
type Login struct {
	nameOrEmail string `json:"nameOrEmail"`
	userPasswd  string `json:"userPasswd"`
}

var (
	MD5PASSWD = ""
	passwd    = ""
)

func init() {
	passwd = conf.GetVal("profile", "password")
	MD5PASSWD = pwmd5(passwd)
	log.Info.Println(MD5PASSWD)
}
func main() {
	//passwd:=conf.GetVal("profile","password")
	log.Info.Println(passwd)
}
func pwmd5(v string) string {
	d := []byte(v)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}
func httppost() {
	//data := `{"type":"10","msg":"hello."}`
	p := new(Login)
	p.nameOrEmail=conf.GetVal("profile","userEmail")
	p.userPasswd=""

	if err := json.Unmarshal([]byte(str), p); err != nil {
		fmt.Println(err)
	}

	request, _ := http.NewRequest("POST", "https://pwl.icu/msg",bytes.NewBuffer([]byte(profiles) ))
	//post数据并接收http响应
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Printf("post data error:%v\n", err)
	} else {
		fmt.Println("post a data successful.")
		respBody, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("response data:%v\n", string(respBody))
		var a Login
		if err = json.Unmarshal(respBody, &a); err != nil {
			fmt.Printf("Unmarshal err, %v\n", err)
			return
		}
		fmt.Printf("%+v", a)
	}
}
