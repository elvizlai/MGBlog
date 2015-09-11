package user

import (
	"gopkg.in/mgo.v2/bson"
	"util"
	"time"
	"model"
	"gopkg.in/mgo.v2"
)

//新建用户
func AddUser(email, nickname, password string) error {
	salt := util.RandString(8)
	password = util.Md5(salt + password)
	u := &User{Id:bson.NewObjectId(), Email:email, NickName:nickname, Salt:salt, Password:password, CreateTime:time.Now()}

	var err error
	model.UserC.Do(func(c *mgo.Collection) {
		err = c.Insert(u)
	})

	return err
}

func GetUserByEmail(email string) *User {
	u := &User{}
	var err error
	model.UserC.Do(func(c *mgo.Collection) {
		err = c.Find(bson.M{"email":email}).One(u)
	})

	if err != nil {
		return nil
	}
	return u
}

func GetUserById(id bson.ObjectId) *User {
	u := new(User)
	model.UserC.Do(func(c *mgo.Collection) {
		if c.FindId(id).One(u) != nil {
			u = nil
		}
	})
	return u
}

func SetToken(email string, token string) {
	model.UserC.Do(func(c *mgo.Collection) {
		c.Update(bson.M{"email":email}, bson.M{"$set":bson.M{"token":token}})
	})
}

//修改密码
func ChangePwd(email, oriPwd, newPwd string) (err error) {
	salt := util.RandString(8)
	pwd := util.Md5(salt + newPwd)
	model.UserC.Do(func(c *mgo.Collection) {
		err = c.Update(bson.M{"email":email}, bson.M{"$set":bson.M{"salt":salt, "password":pwd}})
	})
	return
}