package main

import (
	"fmt"
	"generate/fetcher"
)

func main() {
	conf := fetcher.InitConfigFromJson()
	fmt.Println(conf)
}
