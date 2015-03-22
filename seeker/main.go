package main

import (
	"github.com/astaxie/beego"
	_ "seeker/routers"
)

func main() {
	beego.SetLogger("file", `{"filename":"logs/seeker.log"}`)
	beego.SetLevel(beego.LevelInformational)
	beego.SessionOn = true
	beego.Run()
}
