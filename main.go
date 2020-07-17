package main

import (
	"bytes"
	"fmt"
	"generate/execute"
	"generate/fetcher"
	"os"
	"text/template"
)

func main() {
	conf := fetcher.InitConfigFromJson()
	tems := execute.GetTemplate(conf)
	for _, tem := range tems {
		tmpl, err := template.ParseFiles("input/wxworkDao.php")
		if err != nil {
			panic(err)
		}

		//var content string
		var buf bytes.Buffer

		err = tmpl.Execute(&buf, tem)
		//err = tmpl.Execute(os.Stdin, tem)
		if err != nil {
			panic(err)
		}

		fileName := "output/" + tem.FileName + ".php"
		dstFile, err := os.Create(fileName)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		defer dstFile.Close()
		dstFile.WriteString(buf.String())
	}
}
