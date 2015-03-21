package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplNames = "google.html"
}

func (c *MainController) Post() {
	keyWords := c.GetString("searchKeyWords")
	fmt.Println("keyWords = ", keyWords)
	resp, err := http.Get("https://www.google.com/search?q=" + keyWords + "&ie=UTF-8&oe=UTF-8&hl=en-US")
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

	c.Ctx.WriteString(string(body))
}
