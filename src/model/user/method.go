package user

import (
	"gopkg.in/mgo.v2/bson"
	"util"
	"time"
	"model"
	"gopkg.in/mgo.v2"
)

//create an user, with dup_key error for email or nickname.
func AddUser(email, nickname, password string) error {
	salt := util.RandString(8)
	password = util.Md5(salt + password)
	u := &User{Id:bson.NewObjectId(), Email:email, NickName:nickname, Salt:salt, Password:password, CreateTime:time.Now()}

	var err error
	model.UserC.Do(func(c *mgo.Collection) {
		err = c.Insert(u)
		if err != nil && !mgo.IsDup(err) {
			model.ErrorLog(model.UserC, err, u)
		}
	})

	return err
}

//return nil if user not exist
func GetUserByEmail(email string) *User {
	u := new(User)
	model.UserC.Do(func(c *mgo.Collection) {
		if c.Find(bson.M{"email":email}).One(u) != nil {
			u = nil
		}
	})
	return u
}

//return nil if user not exist
func GetUserById(id bson.ObjectId) *User {
	u := new(User)
	model.UserC.Do(func(c *mgo.Collection) {
		if c.FindId(id).One(u) != nil {
			u = nil
		}
	})
	return u
}

//setting token for login user
func SetToken(id bson.ObjectId, token string) {
	model.UserC.Do(func(c *mgo.Collection) {
		if err := c.UpdateId(id, bson.M{"$set":bson.M{"token":token}}); err != nil {
			model.ErrorLog(model.UserC, err, id.Hex() + ";" + token)
		}
	})
}

//modify password
func ChangePwd(email, newPwd string) (err error) {
	salt := util.RandString(8)
	pwd := util.Md5(salt + newPwd)
	model.UserC.Do(func(c *mgo.Collection) {
		err = c.Update(bson.M{"email":email}, bson.M{"$set":bson.M{"salt":salt, "password":pwd}})
	})
	return
}