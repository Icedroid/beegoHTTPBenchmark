package routers

import (
	"github.com/astaxie/beego"
	"hello/controllers"
)

func init() {
	beego.RESTRouter("/object", &controllers.ObjectController{})
	beego.Router("/set", &controllers.ObjectController{}, "GET:SetRedisHash")
	beego.Router("/get", &controllers.ObjectController{}, "GET:GetRedisHash")
	beego.Router("/log", &controllers.ObjectController{}, "GET:GetAndLog")
}
