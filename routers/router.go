package routers

import (
	"1pay/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/alipay", &controllers.AlipayController{})
	beego.Router("/wepay", &controllers.WepayController{})
}
