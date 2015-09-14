package user

import (
	"model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type User struct {
	Id         bson.ObjectId `bson:"_id"`
	Email      string
	NickName   string
	Salt       string
	Password   string
	Token      string    //token for login
	CreateTime time.Time //user create time
}

func init() {
	model.UserC.Do(func(c *mgo.Collection) {
		c.EnsureIndex(mgo.Index{Key:[]string{"email"}, Unique:true})//邮箱不可重复
		c.EnsureIndex(mgo.Index{Key:[]string{"nickname"}, Unique:true})//昵称不可重复
	})
}