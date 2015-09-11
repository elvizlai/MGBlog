/**
 * Created by elvizlai on 2015/7/28 15:39
 * Copyright © PubCloud
 */

package controller

type LoginController struct {
	BaseController
}

func (this *LoginController) Get() {
	if this.CurrentUser != nil {
		this.Redirect("/newArticle", 302)
		return
	}

	this.TplNames = "login.html"
}