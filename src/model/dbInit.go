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
	UserC collectionName = "Users"//user collection
	ArticleC collectionName = "Articles"//article collection
	FileC collectionName = "Files"//file collection
	StatisticC collectionName = "Statistics" //statistic collection
	TrashC collectionName = "Trashes"//trash collection for deleted articles
)

const (
	historyC collectionName = "History" //to save articles modify history
	errorC collectionName = "Errors"//error log
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

//get a copy of mongo session
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

//modify history
func AddHistory(from collectionName, data interface{}) {
	historyC.Do(func(c *mgo.Collection) {
		c.Insert(bson.M{"from":from, "data":data, "createtime":time.Now()})
	})
}

//add deleted article to trash
func Move2Trash(from collectionName, data interface{}) {
	TrashC.Do(func(c *mgo.Collection) {
		c.Insert(bson.M{"from":from, "data":data, "createtime":time.Now()})
	})
}

//error log
func ErrorLog(from collectionName, err interface{}, data interface{}) {
	errorC.Do(func(c *mgo.Collection) {
		c.Insert(bson.M{"from":from, "err":err, "data":data, "createtime":time.Now()})
	})
}