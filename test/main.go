package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type configuration struct {
	DbType string
	Url    string
}

func main() {
	file, _ := os.Open("conf.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	conf := configuration{}
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(conf)
}
