/**
 * Created by elvizlai on 2015/8/28 11:10
 * Copyright © PubCloud
 */
package console
import (
	"model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"model/static"
	"github.com/astaxie/beego"
)

type VisitorMap struct {
	beego.Controller
}

func (this *VisitorMap) Get() {
	this.TplNames = "visitormap.html"
}

func (this *VisitorMap) Post() {
	visitorList := []static.Static{}
	model.StatisticC.Do(func(c *mgo.Collection) {
		c.Find(bson.M{"infered":true, "city":bson.M{"$ne":""}}).Snapshot().All(&visitorList)
	})

	type Result struct {
		Data     []interface{} `json:"data"`
		GeoCoord map[string][2]float64 `json:"geoCoord"`
	}

	result := new(Result)
	result.GeoCoord = map[string][2]float64{}
	temp := map[string]int{}
	for _, v := range visitorList {
		//先找城市，如果找的到，
		if _, ok := temp[v.City]; ok {
			temp[v.City] += v.PV
		}else {
			temp[v.City] = v.PV
		}
		result.GeoCoord[v.City] = v.Geo
	}

	for k, v := range temp {
		result.Data = append(result.Data, map[string]interface{}{"name":k, "value":v})
	}

	this.Data["json"] = result
	this.ServeJson()
}