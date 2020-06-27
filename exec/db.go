package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const file_url string = "db.json"

type DBConfig struct {
	dbType string
	url    string
	//port    string
	//uname   string
	//passwd  string
	//dbName  string
	//charSet string
}

/**
*
   "dbType" : "mysql",
   "url" :"192.168.240.47",
   "port":"3306",
   "uname":"homework",
   "passwd":"homework",
   "dbName":"homework_wechat_wxsk",
   "charSet": "utf8"
*/

func main() {
	fmt.Println(InitConfigFromJson())
}

func InitConfigFromJson() DBConfig {
	// 打开文件
	file, err := os.Open("db.json")
	if err != nil {
		panic(err)
	}
	// 关闭文件
	defer file.Close()
	//NewDecoder创建一个从file读取并解码json对象的*Decoder，解码器有自己的缓冲，并可能超前读取部分json数据。
	decoder := json.NewDecoder(file)
	//Decode从输入流读取下一个json编码值并保存在v指向的值里
	conf := DBConfig{}
	err = decoder.Decode(&conf)
	if err != nil {
		panic(err)
	}
	fmt.Println(conf.url)
	return conf
}
