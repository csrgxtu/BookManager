package controllers

import (
	"github.com/astaxie/beego"
	"rigo/models"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	var rt models.Results

	this.Data["json"] = &rt
	this.ServeJson()
}
