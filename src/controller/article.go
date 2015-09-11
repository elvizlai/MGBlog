/**
 * Created by elvizlai on 2015/8/24 15:39
 * Copyright © PubCloud
 */
package controller
import (
	"gopkg.in/mgo.v2/bson"
	"model/article"
)

type Article struct {
	BaseController
	art *article.Article
}

func (this *Article) Prepare() {
	this.BaseController.Prepare()

	//找不到该文章
	if this.art = article.GetArticleById(bson.ObjectIdHex(this.Ctx.Input.Param(":id")), true); this.art == nil {
		this.Abort("404")
	}
}

func (this *Article) Get() {
	this.TplNames = "article.html"

	if this.CurrentUser != nil && this.CurrentUser.Id == this.art.UId {
		this.Data["canModify"] = true
	}
	this.Data["article"] = this.art
}