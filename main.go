package main

import (
	"github.com/astaxie/beego"
	_ "hello/docs"
	_ "hello/routers"
)

//		Objects

//	URL					HTTP Verb				Functionality
//	/object				POST					Creating Objects
//	/object/<objectId>	GET						Retrieving Objects
//	/object/<objectId>	PUT						Updating Objects
//	/object				GET						Queries
//	/object/<objectId>	DELETE					Deleting Objects

func main() {
	//	if beego.RunMode == "dev" {
	//		beego.DirectoryIndex = true
	//		beego.StaticDir["/swagger"] = "swagger"
	//	}

	beego.Run()
}
