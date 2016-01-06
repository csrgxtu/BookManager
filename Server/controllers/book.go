package controllers

import (
	"github.com/astaxie/beego"
)

type BookController struct {
	beego.Controller
}

func (this *BookController) Create() {
  c.TplNames = "index.tpl"
}
