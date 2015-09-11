/**
 * Created by elvizlai on 2015/8/24 13:44
 * Copyright © PubCloud
 */
package console
import (
	"enum"
	"controller"
)

//需要鉴权的
type AuthController struct {
	controller.BaseController
}

func (this *AuthController) Prepare() {
	this.BaseController.Prepare()

	//未登录
	if this.CurrentUser == nil {
		switch this.Ctx.Request.Method {
		case "POST":
			this.CustomAbort(enum.UnAuthorized.Code(), enum.UnAuthorized.Str())
		default:
			this.Redirect("/login", 302)
			this.StopRun()
		}
	}
}

