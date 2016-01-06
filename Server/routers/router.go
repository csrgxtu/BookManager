package routers

import (
	"BookManager/Server/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

		// BookController
		// beego.Router("/article", &controllers.MainController{}, "post:NewArticle")
		beego.Router("/book", &controllers.BookController{}, "put:CreateBook")
		beego.Router("/book", &controllers.BookController{}, "get:ReadBook")
		beego.Router("/book", &controllers.BookController{}, "post:UpdateBook")
		beego.Router("/book", &controllers.BookController{}, "delete:DeleteBook")

		beego.Router("/books", &controllers.BookController{}, "put:CreateBooks")
		beego.Router("/books", &controllers.BookController{}, "get:ReadBooks")
		beego.Router("/books", &controllers.BookController{}, "post:UpdateBooks")
		beego.Router("/books", &controllers.BookController{}, "delete:DeleteBooks")
}
