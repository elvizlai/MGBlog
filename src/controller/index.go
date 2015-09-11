/**
 * Created by elvizlai on 2015/7/28 16:46
 * Copyright © PubCloud
 */

package controller
import (
	"github.com/astaxie/beego/utils/pagination"
	"model"
	"model/article"
	"gopkg.in/mgo.v2"
	"strconv"
)

const limit = 10

type Index struct {
	BaseController
}

//文章列表
func (this *Index) Get() {
	this.TplNames = "index.html"

	currentPage := 1
	if page, err := strconv.Atoi(this.Input().Get("p")); err == nil {
		currentPage = page
	}

	articles := []article.Article{}
	totalCount := 1
	model.ArticleC.Do(func(c *mgo.Collection) {
		q := c.Find(nil).Sort("-createtime")
		totalCount, _ = q.Count()
		q.Skip(limit * (currentPage - 1)).Limit(limit).All(&articles)
	})

	this.Data["articles"] = articles
	pagination.SetPaginator(this.Ctx, limit, int64(totalCount))
}