package controllers

import (
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
	beego.Info("[##Data##]Search key Words: ", keyWords)
	targetUrl := "https://www.google.com/search?q=" + keyWords + "&ie=UTF-8&oe=UTF-8&hl=en-US"
	resp, err := http.Get(targetUrl)
	if err != nil {
		beego.Error("Fail to execute http.Get(", targetUrl, ").")
		return
	}
	defer resp.Body.Close()

	isOpenEyes := false
	isOpenStr := c.GetString("isOpenEyes")

	if "on" == isOpenStr {
		isOpenEyes = true
	} else {
		isOpenEyes = false
	}
	beego.Info("[##Data##]Is open eyes:", isOpenEyes)
	c.SetSession("openEyeSession", isOpenEyes)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		beego.Error("Fail to read response's body!")
		return
	}

	c.Ctx.WriteString(string(body))
}

type RedirectController struct {
	beego.Controller
}

func (c *RedirectController) Get() {
	targetUrl := c.GetString("q")
	beego.Info("[##Data##]The redirect target url is ", targetUrl)
	isOpenEyes := false
	openEyeSession := c.GetSession("openEyeSession")
	if openEyeSession != nil {
		isOpenEyes, _ = openEyeSession.(bool)
		//if err != nil {
		//	isOpenEyes = false
		//}
	}
	if isOpenEyes {
		beego.Info("Visit site[", targetUrl, "] by proxy server!")
		resp, err := http.Get(targetUrl)
		if err != nil {
			beego.Error("Fail to execute http.Get(", targetUrl, ").")
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			beego.Error("Fail to read response's body!")
			return
		}

		c.Ctx.WriteString(string(body))
	} else {
		beego.Info("Visit site[", targetUrl, "] by local redirection!")
		c.Ctx.Redirect(http.StatusFound, targetUrl)
	}
}
