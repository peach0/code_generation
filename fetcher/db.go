package fetcher

import (
	"encoding/json"
	"os"
)

const file_url string = "conf/db.json"

type DBConfig struct {
	DbType  string
	Url     string
	Port    string
	Uname   string
	Passwd  string
	DbName  string
	CharSet string
	Table   []string
}

func InitConfigFromJson() DBConfig {
	// 打开文件
	file, err := os.Open(file_url)
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
	return conf
}
