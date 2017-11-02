package controllers

import (
	"github.com/astaxie/beego"
)

type AlipayController struct {
	beego.Controller
}

func (this *AlipayController) Get() {
	this.TplName = "alipay.html"
}
