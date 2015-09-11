package router

import (
	"github.com/astaxie/beego"
	"controller/console"
	"controller"
)

func init() {
	beego.Router("/", &controller.Index{})
	beego.Router("/archives", &controller.Archives{})
	beego.Router("/me", &controller.Me{})


	beego.Router("/article/:id([0-9a-z]{24})", &controller.Article{})
	beego.Router("/file/:id([0-9a-z]{24})", &controller.FileController{})
	beego.Router("/tag/:tag", &controller.Tag{})

	beego.Router("/login", &controller.LoginController{})
	beego.Router("/register", &controller.RegisterController{})
	beego.Router("/newArticle", &console.NewArticle{})
	beego.Router("/upload", &console.Upload{})
	beego.Router("/modifyArticle/:id([0-9a-z]{24})", &console.ModifyArticle{})
	beego.Router("/trash/?:id([0-9a-z]{24})", &console.Trash{})//回收站
	beego.Router("/visitormap", &console.VisitorMap{})
}