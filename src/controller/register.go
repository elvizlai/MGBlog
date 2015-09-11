/**
 * Created by elvizlai on 2015/7/28 18:06
 * Copyright © PubCloud
 */
package controller
import (
	"github.com/astaxie/beego"
)

type RegisterController struct {
	BaseController
}

func (this *RegisterController) Get() {
	this.TplNames = "register.html"

	//是否允许注册
	if cr, ok := beego.AppConfig.Bool("CanBeRegister"); ok == nil {
		this.Data["CanBeRegister"] = cr
	}
}