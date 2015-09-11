/**
 * Created by elvizlai on 2015/7/28 15:39
 * Copyright © PubCloud
 */
package controller
import (
	"github.com/astaxie/beego"
	"github.com/bitly/go-simplejson"
	"enum"
	"model/user"
	"gopkg.in/mgo.v2/bson"
	"model/static"
)

//不需要鉴权的
type BaseController struct {
	beego.Controller
	CurrentUser *user.User //当前登入账户
}

//是否登录
func (this *BaseController) Prepare() {
	if this.Ctx.Input.Method() == "GET" {
		this.Layout = "Layout.html"
		this.Data["visitCount"] = static.GetVisitCount()//todo easy
	}

	if uId := this.GetSession("uId"); uId != nil {
		if this.CurrentUser = user.GetUserById(bson.ObjectIdHex(uId.(string))); this.CurrentUser.Token == this.StartSession().SessionID() {
			this.Data["isLogin"] = true
		}else {
			this.CurrentUser = nil
		}
	}
}

func (this *BaseController) ReqJson() *simplejson.Json {
	defer func() {
		if err := recover(); err != nil&&err.(string) == "jsonErr" {
			this.CustomAbort(enum.BadRequest.Code(), enum.BadRequest.Str())
		}
	}()

	if json, err := simplejson.NewJson(this.Ctx.Input.RequestBody); err == nil {
		return json
	}else {
		panic("jsonErr")
	}
}

func (this *BaseController) RespJson(e enum.Code, result interface{}) {
	this.Data["json"] = map[string]interface{}{"code":e.Code(), "msg":e.Str(), "result":result}
	this.ServeJson()
	//this.StopRun()
}