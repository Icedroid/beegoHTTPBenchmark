package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["hello/controllers:CMSController"] = append(beego.GlobalControllerRouter["hello/controllers:CMSController"],
		beego.ControllerComments{
			"StaticBlock",
			"[get]",
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["hello/controllers:CMSController"] = append(beego.GlobalControllerRouter["hello/controllers:CMSController"],
		beego.ControllerComments{
			"Product",
			"/products",
			[]string{"get"},
			nil})

}
