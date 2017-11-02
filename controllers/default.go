package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/astaxie/beego"
	// "net/http"
)

type MainController struct {
	beego.Controller
	http.Request
}

func (this *MainController) Get() {
	UserAgent := this.Ctx.Request.UserAgent()
	fmt.Println(UserAgent)
	// fmt.Println("UA:", strings.Contains(UserAgent, "MicroMessenger"))
	this.Data["Website"] = "codebeta.cn"

	// c.Data["Email"] = "astaxie@gmail.com"
	this.TplName = "index.html"
	if strings.Contains(UserAgent, "MicroMessenger") {
		this.Redirect("/wepay", 302)
	} else if strings.Contains(UserAgent, "AlipayClient") {
		this.Redirect("HTTPS://QR.ALIPAY.COM/FKX041661UPPDDXQ4FQW36", 302)
	} else {
		this.Data["App"] = "请使用微信或者支付宝客户端扫码" //this.Ctx.Request.UserAgent()
	}
}
