package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.ConBtroller
}

func (c *MainController) Get() {
	c.TplNames = "google.html"
}
