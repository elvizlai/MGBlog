package file
import (
	"gopkg.in/mgo.v2/bson"
	"time"
	"model"
	"gopkg.in/mgo.v2"
)

type File struct {
	Id         bson.ObjectId `bson:"_id"`
	Name       string        //文件名
	MD5        string        //文件的md5值
	Data       []byte        //数据
	CreateTime time.Time     //上传时间
	UId        bson.ObjectId //作者
}

func init() {
	model.FileC.Do(func(c *mgo.Collection) {
		//md5唯一，可能会撞库，虽然几率超级小...//todo
		c.EnsureIndex(mgo.Index{Key:[]string{"md5"}, Unique:true, DropDups:true})
	})
}