package controllers

import (
	"github.com/astaxie/beego"
)

type WepayController struct {
	beego.Controller
}

func (this *WepayController) Get() {
	this.TplName = "weixin.html"
}
