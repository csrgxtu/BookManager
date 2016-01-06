package main

import (
	_ "BookManager/Server/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

