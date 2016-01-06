package routers

import (
	"BookManager/Server/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

		// BookController
		beego.Router("/book", &controllers.BookController{}, "Put: CreateBook")
		beego.Router("/book", &controllers.BookController{}, "Get: ReadBook")
		beego.Router("/book", &controllers.BookController{}, "Post: UpdateBook")
		beego.Router("/book", &controllers.BookController{}, "Delete: DeleteBook")

		beego.Router("/books", &controllers.BookController{}, "Put: CreateBooks")
		beego.Router("/books", &controllers.BookController{}, "Get: ReadBooks")
		beego.Router("/books", &controllers.BookController{}, "Post: UpdateBooks")
		beego.Router("/books", &controllers.BookController{}, "Delete: DeleteBooks")

}
