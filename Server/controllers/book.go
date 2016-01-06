package controllers

import (
	"github.com/astaxie/beego"
  "BookManager/Server/models"
	"BookManager/Server/services"
  "encoding/json"
)

type BookController struct {
	beego.Controller
}

func (this *BookController) CreateBook() {
  var book models.Book
  json.Unmarshal(this.Ctx.Input.RequestBody, &book)

  var rt models.Results

  nBook, err := services.CreateBook(book)
  if err != nil {
    rt.Error = true
    rt.Msg = err.Error()
  } else {
    rt.Error = false
    rt.Msg = "Successful"
    rt.Total = 1
    rt.Data = make([]models.Recs, rt.Total)
    rt.Data[0] = nBook
  }

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
