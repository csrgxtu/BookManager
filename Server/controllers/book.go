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

func (this *BookController) Read() {
  var rt models.Results

  this.Data["json"] = &rt
  this.ServeJson()
}

func (this *BookController) Update() {
  var rt models.Results

  this.Data["json"] = &rt
  this.ServeJson()
}

func (this *BookController) Delete() {
  var rt models.Results

  this.Data["json"] = &rt
  this.ServeJson()
}
