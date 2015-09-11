/**
 * Created by elvizlai on 2015/8/24 14:33
 * Copyright © PubCloud
 */
package article
import (
	"gopkg.in/mgo.v2/bson"
	"model"
	"gopkg.in/mgo.v2"
	"time"
	"model/user"
	"regexp"
)

//通过id获取文章，static为是否统计的标识
func GetArticleById(id bson.ObjectId, static ...bool) *Article {
	art := new(Article)
	model.ArticleC.Do(func(c *mgo.Collection) {
		if c.FindId(id).One(art) != nil {
			art = nil
		}else {
			//previous
			previous := new(Article)
			if c.Find(bson.M{"createtime":bson.M{"$gt":art.CreateTime}}).Sort("createtime").One(previous) == nil {
				art.Previous = previous
			}
			//next
			next := new(Article)
			if c.Find(bson.M{"createtime":bson.M{"$lt":art.CreateTime}}).Sort("-createtime").One(next) == nil {
				art.Next = next
			}

			if u := user.GetUserById(art.UId); u != nil {
				art.AuthStr = u.NickName
			}

			if static != nil {
				c.UpdateId(id, bson.M{"$inc":bson.M{"pv": 1}})
			}
		}
	})
	return art
}

//文章发布
func AddArticle(uId bson.ObjectId, title string, tags []string, markdown, htmlContent string) error {
	var err error
	model.ArticleC.Do(func(c *mgo.Collection) {
		t := time.Now()
		art := &Article{Id:bson.NewObjectId(), UId:uId, Title:title, Tags:tags, Abstract:getAbstract(htmlContent), Markdown:markdown, HtmlContent:htmlContent, PV:1, CreateTime:t, UpdateTime:t}
		err = c.Insert(art)
	})
	return err
}

//文章修改
func ModifyArticle(id bson.ObjectId, title string, tags []string, markdown, htmlContent string) error {
	var err error
	//先将上一历史存档
	model.ArticleC.Do(func(c *mgo.Collection) {
		art := new(Article)
		if err = c.FindId(id).One(art); err == nil {
			model.AddHistory(model.ArticleC, art)
		}else {
			model.ErrorLog(model.ArticleC, err, art)
		}

		art.Title = title
		art.Tags = tags
		art.Abstract = getAbstract(htmlContent)
		art.HtmlContent = htmlContent
		art.Markdown = markdown
		art.UpdateTime = time.Now()
		if err = c.UpdateId(id, art); err != nil {
			model.ErrorLog(model.ArticleC, err, art)
		}
	})
	return err
}

//文章删除
func DeleteArticle(id bson.ObjectId) error {
	var err error
	model.ArticleC.Do(func(c *mgo.Collection) {
		//先存档
		art := new(Article)
		if err = c.FindId(id).One(art); err == nil {
			art.UpdateTime = time.Now()
			model.Move2Trash(model.ArticleC, art)
		}else {
			model.ErrorLog(model.ArticleC, err, art)
		}

		if err = c.RemoveId(id); err != nil {
			model.ErrorLog(model.ArticleC, err, art)
		}
	})
	return err
}

func getAbstract(markdown string) string {
	reg := regexp.MustCompile(`<!-{2,}more-{2,}>`)
	index := reg.FindStringIndex(markdown)
	abstract := ""
	if index != nil {
		abstract = markdown[:index[0]]
	}
	return abstract
}