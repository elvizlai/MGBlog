package file
import (
	"gopkg.in/mgo.v2/bson"
	"model"
	"gopkg.in/mgo.v2"
	"time"
	"model/user"
	"os"
	"io/ioutil"
	"crypto/md5"
	"encoding/hex"
)

func AddFile(name string, data []byte, usr *user.User) bson.ObjectId {
	//计算md5
	md5h := md5.New()
	md5h.Write(data)
	m := hex.EncodeToString(md5h.Sum([]byte("")))

	f := new(File)
	model.FileC.Do(func(c *mgo.Collection) {
		//先查询，如果找的到，就直接返回该图片
		if c.Find(bson.M{"md5":m}).One(f) != nil {
			f.Id = bson.NewObjectId()
			f.Name = name
			f.Data = data
			f.MD5 = m
			f.CreateTime = time.Now()
			f.UId = usr.Id
			c.Insert(f)
			//backup uploaded file
			go func(u *user.User, d []byte) {
				os.MkdirAll("uploads/" + u.Email, os.ModePerm)
				ioutil.WriteFile("uploads/" + u.Email + "/" + name, d, 0660)
			}(usr, data)
		}
	})

	return f.Id
}

func GetFileById(id bson.ObjectId) *File {
	file := new(File)
	model.FileC.Do(func(c *mgo.Collection) {
		if c.FindId(id).One(file) != nil {
			file = nil
		}
	})
	return file
}