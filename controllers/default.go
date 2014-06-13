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
	objectid := models.AddOneToRedis(ob)
	this.Data["json"] = map[string]string{"ObjectId": objectid}
	this.ServeJson()
}

func (this *ObjectController) GetRedisHash() {
	ob, err := models.GetObject()
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = ob
	}
	this.ServeJson()
}