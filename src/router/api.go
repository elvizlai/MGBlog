package router

import (
	"github.com/astaxie/beego"
	"controller/api"
)

func init() {
	beego.Router("/api/register", &api.User{}, "post:Register")
	beego.Router("/api/login", &api.User{}, "post:Login")
}