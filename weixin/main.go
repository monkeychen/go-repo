package main

import (
	"github.com/astaxie/beego"
	_ "weixin/routers"
)

func main() {
	beego.SetLogger("file", `{"filename":"logs/weixin.log"}`)
	beego.SetLevel(beego.LevelInformational)
	beego.SessionOn = true
	beego.Run()
}
