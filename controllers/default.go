package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"hello/models"
	"log"
)

type ResponseInfo struct {
}

type ObjectController struct {
	beego.Controller
}

func (this *ObjectController) Post() {
	var ob models.Object
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	objectid := models.AddOne(ob)
	this.Data["json"] = map[string]string{"ObjectId": objectid}
	this.ServeJson()
}

func (this *ObjectController) Get() {
	objectId := this.Ctx.Input.Params[":objectId"]
	if objectId != "" {
		ob, err := models.GetOne(objectId)
		if err != nil {
			this.Data["json"] = err
		} else {
			this.Data["json"] = ob
		}
	} else {
		obs := models.GetAll()
		this.Data["json"] = obs
	}
	this.ServeJson()
}

func (this *ObjectController) Put() {
	objectId := this.Ctx.Input.Params[":objectId"]
	var ob models.Object
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)

	err := models.Update(objectId, ob.Score)
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = "update success!"
	}
	this.ServeJson()
}

func (this *ObjectController) Delete() {
	objectId := this.Ctx.Input.Params[":objectId"]
	models.Delete(objectId)
	this.Data["json"] = "delete success!"
	this.ServeJson()
}

func (this *ObjectController) GetAndLog() {
	obs := models.GetAll()
	log.Printf("objects=%v", obs)
	this.Data["json"] = obs
	this.ServeJson()
}

func (this *ObjectController) SetRedisHash() {
	var ob models.Object
	ob.PlayerName = "Icecroid"
	ob.Score = 100
	objectid := models.RAddOne(ob)
	this.Data["json"] = map[string]string{"ObjectId": objectid}
	this.ServeJson()
}

func (this *ObjectController) GetRedisHash() {
	ob, err := models.RGetObject()
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = ob
	}
	this.ServeJson()
}

func (this *ObjectController) GetRedisHashV2() {
	ob, err := models.RGetObject2()
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = ob
	}
	this.ServeJson()
}

func (this *ObjectController) GetRedisHashV3() {
	ob, err := models.RGetObject3()
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = ob
	}
	this.ServeJson()
}

func (this *ObjectController) AddMongoRow() {
	var ob models.Object
	ob.PlayerName = "Icecroid"
	ob.Score = 100
	objectid := models.MAddOne(ob)
	this.Data["json"] = map[string]string{"ObjectId": objectid}
	this.ServeJson()
}

func (this *ObjectController) GetMongoData() {
	objectId := this.Ctx.Input.Params[":objectId"]
	if objectId != "" {
		ob, err := models.MGetOne(objectId)
		if err != nil {
			this.Data["json"] = err
		} else {
			this.Data["json"] = ob
		}
	} else {
		obs, err := models.MGetObject()
		if err != nil {
			this.Data["json"] = err
		} else {
			this.Data["json"] = obs
		}
	}
	this.ServeJson()
}

func (this *ObjectController) GetMongoData2() {
	objectId := this.Ctx.Input.Params[":objectId"]
	if objectId != "" {
		ob, err := models.MGetOne2(objectId)
		if err != nil {
			this.Data["json"] = err
		} else {
			this.Data["json"] = ob
		}
	} else {
		obs, err := models.MGetObject2()
		if err != nil {
			this.Data["json"] = err
		} else {
			this.Data["json"] = obs
		}
	}
	this.ServeJson()
}
