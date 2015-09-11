/**
 * Created by elvizlai on 2015/8/27 09:16
 * Copyright © PubCloud
 */
package static
import (
	"model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
	"util"
)

func init() {
	//定时更新任务
	go func() {
		for {
			model.StatisticC.Do(func(c *mgo.Collection) {
				iter := c.Find(bson.M{"infered":bson.M{"$ne":true}}).Snapshot().Iter()
				s := new(Static)
				for iter.Next(s) {
					if result := util.InfoGeoByIP(s.IP); result != nil {
						//添加返回值
						s.City = result.Get("content").Get("address_detail").Get("city").MustString()
						s.Geo[0] = util.Str2Float(result.Get("content").Get("point").Get("x").MustString())
						s.Geo[1] = util.Str2Float(result.Get("content").Get("point").Get("y").MustString())
					}
					s.Infered = true
					c.Update(bson.M{"ip":s.IP}, s)
				}
			})
			<-time.After(time.Hour)
		}
	}()

}

func Add(ip string, url string) {
	model.StatisticC.Do(func(c *mgo.Collection) {
		c.Upsert(bson.M{"ip":ip}, bson.M{"$inc":bson.M{"pv": 1}, "$addToSet":bson.M{"path":url}, "$set":bson.M{"lastvisit":time.Now()}})
	})
}

func GetVisitCount() (count int) {
	model.StatisticC.Do(func(c *mgo.Collection) {
		count, _ = c.Count()
	})
	return
}