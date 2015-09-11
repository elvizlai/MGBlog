/**
 * Created by elvizlai on 2015/8/25 16:41
 * Copyright © PubCloud
 */
package controller
import (
	"model"
	"gopkg.in/mgo.v2"
	"model/article"
	"github.com/astaxie/beego/utils/pagination"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

type Tag struct {
	BaseController
}

func (this *Tag) Get() {
	tag := this.Ctx.Input.Param(":tag")

	if tag == "" {
		this.Redirect("/", 302)
	}

	this.TplNames = "tag.html"

	currentPage := 1
	if page, err := strconv.Atoi(this.Input().Get("p")); err == nil {
		currentPage = page
	}

	totalCount := 1
	articles := []article.Article{}
	model.ArticleC.Do(func(c *mgo.Collection) {
		q := c.Find(bson.M{"tags":tag}).Sort("-createtime")
		totalCount, _ = q.Count()
		q.Skip(limit * (currentPage - 1)).Limit(limit).All(&articles)
	})

	this.Data["articles"] = articles
	this.Data["tag"] = "Tag:" + tag

	pagination.SetPaginator(this.Ctx, limit, int64(totalCount))
}