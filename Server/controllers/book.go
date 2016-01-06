package controllers

import (
	"github.com/astaxie/beego"
  "rigo/models"
)

type BookController struct {
	beego.Controller
}

func (this *BookController) Create() {
  var rt models.Results

  this.Data["json"] = &rt
  this.ServeJson()
}
