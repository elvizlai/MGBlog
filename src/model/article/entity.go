/**
 * Created by elvizlai on 2015/8/24 14:26
 * Copyright © PubCloud
 */
package article
import (
	"gopkg.in/mgo.v2/bson"
	"time"
	"model"
	"gopkg.in/mgo.v2"
)

type Article struct {
	Id          bson.ObjectId `bson:"_id"`
	UId         bson.ObjectId
	AuthStr     string
	Title       string   //标题
	Tags        []string //标签
	Abstract    string   //摘要
	HtmlContent string   //html正文
	Markdown    string   //markdown正文
	PV          int      //浏览量
	CreateTime  time.Time
	UpdateTime  time.Time
	Previous    *Article
	Next        *Article
}

func init() {
	model.ArticleC.Do(func(c *mgo.Collection) {
		c.EnsureIndexKey("tags", "createtime")
	})
}