package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplNames = "google.html"
}

func (c *MainController) Post() {
	c.TplNames = "google.html"

}
