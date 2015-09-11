/**
 * Created by elvizlai on 2015/8/26 16:33
 * Copyright © PubCloud
 */
package console
import (
	"model/article"
	"gopkg.in/mgo.v2/bson"
	"model"
	"gopkg.in/mgo.v2"
	"enum"
	"time"
)
//todo
type Trash struct {
	AuthController
	id bson.ObjectId
}

func (this *Trash) Prepare() {
	this.AuthController.Prepare()

	if idStr := this.Ctx.Input.Param(":id"); idStr != "" {
		this.id = bson.ObjectIdHex(idStr)
	}
}

func (this *Trash) Get() {
	if this.id.Valid() {
		this.TplNames = "article.html"
		art := struct {Data article.Article}{}
		model.TrashC.Do(func(c *mgo.Collection) {
			c.Find(bson.M{"from":model.ArticleC, "data._id":this.id}).Select(bson.M{"data": 1, "_id":0}).One(&art)
		})
		this.Data["article"] = art.Data
	}else {
		this.TplNames = "console/deleted_articles.html"
		result := []struct {Data article.Article}{}
		model.TrashC.Do(func(c *mgo.Collection) {
			c.Find(bson.M{"from":model.ArticleC}).Sort("-data.updatetime").Select(bson.M{"data": 1, "_id":0}).All(&result)
		})
		this.Data["articles"] = result
	}
}

func (this *Trash) Post() {
	if this.id.Valid() {
		art := struct {
			Id   bson.ObjectId `bson:"_id"`
			Data article.Article
		}{}

		model.TrashC.Do(func(c *mgo.Collection) {
			if c.Find(bson.M{"from":model.ArticleC, "data._id":this.id}).Select(bson.M{"data": 1}).One(&art) == nil {
				//从trash中移除要删除的文章
				c.RemoveId(art.Id)

				//文章恢复
				model.ArticleC.Do(func(c *mgo.Collection) {
					t := time.Now()
					art.Data.CreateTime = t
					art.Data.UpdateTime = t
					c.Insert(&art.Data)
				})

				this.RespJson(enum.OK, nil)
			}else {
				this.RespJson(enum.NotFound, nil)
			}
		})
	}
	this.RespJson(enum.BadRequest, nil)
}