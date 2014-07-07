// @APIVersion 1.0.0
// @Title mobile API
// @Description mobile has every tool to get any job done, so codename for the new mobile APIs.
// @Contact astaxie@gmail.com
package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"hello/controllers"
)

func init() {
	var filterCORS = func(ctx *context.Context) {
		ctx.Output.Header("Access-Control-Allow-Origin", "*")
		ctx.Output.Header("Access-Control-Allow-Credentials", "true")
	}
	beego.InsertFilter("/*", beego.BeforeRouter, filterCORS)

	ns :=
		beego.NewNamespace("/v1",
			beego.NSNamespace("/cms",
				beego.NSInclude(
					&controllers.CMSController{},
				),
			),
		)
	beego.AddNamespace(ns)

	//	beego.RESTRouter("/object", &controllers.ObjectController{})
	//	beego.Router("/log", &controllers.ObjectController{}, "GET:GetAndLog")
	//	beego.Router("/rset", &controllers.ObjectController{}, "GET:SetRedisHash")
	//	beego.Router("/rget", &controllers.ObjectController{}, "GET:GetRedisHash")
	//	beego.Router("/rget2", &controllers.ObjectController{}, "GET:GetRedisHashV2")
	//	beego.Router("/rget3", &controllers.ObjectController{}, "GET:GetRedisHashV3")
	//	beego.Router("/mset", &controllers.ObjectController{}, "GET:AddMongoRow")
	//	beego.Router("/mget", &controllers.ObjectController{}, "GET:GetMongoData")
	//	beego.Router("/mget2", &controllers.ObjectController{}, "GET:GetMongoData")
	//	beego.Router("/mget/:objectId", &controllers.ObjectController{}, "GET:GetMongoData")
	//	beego.Router("/mget2/:objectId", &controllers.ObjectController{}, "GET:GetMongoData")

}
