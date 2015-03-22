// template.go
package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"testing"
)

type Person struct {
	UserName string
	email    string //未导出的字段，首字母是小写的
}

func TestMain(t *testing.T) {
	file, _ := os.Create("logs.log")
	logger := log.New(file, "", log.LstdFlags|log.Llongfile)
	tmpl := template.New("fieldname example")
	tmpl, _ = tmpl.Parse("{{if 1}}hello {{.UserName}}!{{else}}else block...{{end}}")
	p := Person{UserName: "Astaxie"}
	logger.Println("begin execute...")
	tmpl.Execute(os.Stdout, p)
	logger.Println("after execute...")
	//httpGet()

}

func httpGet() {
	resp, err := http.Get("https://www.google.com/search?q=golang&ie=UTF-8&oe=UTF-8&hl=en-US")
	if err != nil {
		// handle error
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		fmt.Println(err)
	}

	fmt.Println(string(body))
}

func httpPostForm() {
	resp, err := http.PostForm("https://www.google.com.hk/search",
		url.Values{"q": {"golang"}, "ie": {"UTF-8"}, "oe": {"UTF-8"}, "hl": {"zh-CN"}})

	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))

}
