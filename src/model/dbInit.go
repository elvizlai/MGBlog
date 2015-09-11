package model

import (
	"gopkg.in/mgo.v2"
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
	"time"
)

var session *mgo.Session

const dbName = "MBlog"

type collectionName string

const (
	UserC collectionName = "Users"//用户
	ArticleC collectionName = "Articles"//文章
	FileC collectionName = "Files"//文件
	StatisticC collectionName = "Statistics" //统计分析
	TrashC collectionName = "Trashes"//回收站
//FavorC collectionName = "Favors"//收藏
)

const (
	historyC collectionName = "History" //历史
	errorC collectionName = "Errors"//错误日志

)

func (this collectionName) String() string {
	return string(this)
}

func init() {
	var err error
	session, err = mgo.Dial(beego.AppConfig.String("DBUrl"))
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
}

//获取session--注意session要及时关闭
func getSession() *mgo.Session {
	return session.Copy()
}

func (this collectionName) Do(f func(c *mgo.Collection)) {
	databaseDo(dbName, func(db *mgo.Database) {
		f(db.C(this.String()))
	})
}

func databaseDo(dbName string, f func(db *mgo.Database)) {
	s := getSession()
	defer s.Close()
	f(s.DB(dbName))
}

func AddHistory(from collectionName, data interface{}) {
	historyC.Do(func(c *mgo.Collection) {
		c.Insert(bson.M{"from":from, "data":data, "createtime":time.Now()})
	})
}

func Move2Trash(from collectionName, data interface{}) {
	TrashC.Do(func(c *mgo.Collection) {
		c.Insert(bson.M{"from":from, "data":data, "createtime":time.Now()})
	})
}

func ErrorLog(from collectionName, err interface{}, data interface{}) {
	errorC.Do(func(c *mgo.Collection) {
		c.Insert(bson.M{"from":from, "err":err, "data":data, "createtime":time.Now()})
	})
}