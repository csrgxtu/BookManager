package controllers

import (
	"github.com/astaxie/beego"
  "rigo/models"
)

type BookController struct {
	beego.Controller
}

func (this *BookController) CreateBook() {
  var rt models.Results

  this.Data["json"] = &rt
  this.ServeJson()
}

func (this *BookController) CreateBooks() {
  var rt models.Results

  this.Data["json"] = &rt
  this.ServeJson()
}

func (this *BookController) ReadBook() {
  var rt models.Results

  this.Data["json"] = &rt
  this.ServeJson()
}

func (this *BookController) ReadBooks() {
  var rt models.Results

  this.Data["json"] = &rt
  this.ServeJson()
}

func (this *BookController) UpdateBook() {
  var rt models.Results

  this.Data["json"] = &rt
  this.ServeJson()
}

func (this *BookController) UpdateBooks() {
  var rt models.Results

  this.Data["json"] = &rt
  this.ServeJson()
}

func (this *BookController) DeleteBook() {
  var rt models.Results

  this.Data["json"] = &rt
  this.ServeJson()
}

func (this *BookController) DeleteBooks() {
  var rt models.Results

  this.Data["json"] = &rt
  this.ServeJson()
}
