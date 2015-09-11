/**
 * Created by elvizlai on 2015/8/26 12:59
 * Copyright © PubCloud
 */
package console
import (
	"model/article"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"enum"
)

type ModifyArticle struct {
	AuthController
	art *article.Article
}

func (this *ModifyArticle) Prepare() {
	this.AuthController.Prepare()

	id := bson.ObjectIdHex(this.Ctx.Input.Param(":id"))

	this.art = article.GetArticleById(id)
	if this.art == nil {
		this.Abort("404")//找不到对应的文档
	}
}

func (this *ModifyArticle) Get() {
	this.TplNames = "console/modify_article.html"
	this.Data["article"] = this.art
}

func (this *ModifyArticle) Post() {
	req := this.ReqJson()

	method := req.Get("method").MustString()

	var err error

	if method == "update" {
		title := req.Get("title").MustString()
		tagStr := req.Get("tags").MustString()

		var tags []string
		if tagStr == "" {
			tags = []string{"未分类"}
		}else if strings.Contains(tagStr, ";") {
			tags = strings.Split(tagStr, ";")
		}else {
			tags = strings.Split(tagStr, "；")
		}

		for i, _ := range tags {
			tags[i] = strings.TrimLeft(tags[i], " ")
			tags[i] = strings.TrimRight(tags[i], " ")
		}

		markdown := req.Get("markdown").MustString()
		htmlContent := req.Get("htmlContent").MustString()

		err = article.ModifyArticle(this.art.Id, title, tags, markdown, htmlContent)
	}else if method == "delete" {
		err = article.DeleteArticle(this.art.Id)
	}else {
		this.RespJson(enum.BadRequest, nil)
	}

	if err == nil {
		this.RespJson(enum.OK, nil)
	}else {
		this.RespJson(enum.UNKNOWN, err)
	}
}