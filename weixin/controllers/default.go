package controllers

import (
	"crypto/sha1"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"log"
	"sort"
	"strings"
)

const (
	token = "simiamweixin4go"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	signature := c.GetString("signature")
	timestamp := c.GetString("timestamp")
	nonce := c.GetString("nonce")
	echostr := c.GetString("echostr")

	mySignature := makeSignature(timestamp, nonce)
	if mySignature == signature {
		log.Println("Simiam weixin Service: validate successful!")
		beego.Info("Simiam weixin Service: validate successful!")
		c.Ctx.WriteString(string(echostr))
	} else {
		log.Fatalln("Simiam weixin Service: validate fail!")
		beego.Error("Simiam weixin Service: validate fail!")
	}
	return
}

func makeSignature(timestamp, nonce string) string {
	sl := []string{token, timestamp, nonce}
	sort.Strings(sl)
	s := sha1.New()
	io.WriteString(s, strings.Join(sl, ""))
	return fmt.Sprintf("%x", s.Sum(nil))
}
