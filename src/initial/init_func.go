/**
 * Created by elvizlai on 2015/7/28 15:37
 * Copyright © PubCloud
 */

package initial

import (
	"github.com/astaxie/beego"
	"html/template"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"time"
	"fmt"
)

func init() {
	beego.AddFuncMap("obj2str", obj2str)
	beego.AddFuncMap("slice2str", slice2str)
	beego.AddFuncMap("compare", compare)
	beego.AddFuncMap("set", set)
}

//Usage:将ObjectId转换为String
//{{obj2str .ObjectId}}
func obj2str(obj bson.ObjectId) template.JS {
	return template.JS(obj.Hex())
}

//add by ElvizLai
//Usage:当前模板上下文中设置一个变量
//{{set . "var" "Mes"}}
//{{.var}}
func set(renderArgs map[interface{}]interface{}, key string, value interface{}) template.JS {
	renderArgs[key] = value
	return template.JS("")
}

func slice2str(slice []string) template.JS {
	return template.JS(strings.Join(slice, ";"))
}

func compare(t1, t2 time.Time) template.JS {
	result := 0
	switch  {
	case t1.Before(t2):
		result = -1
	case t1.After(t2):
		result = 1
	}
	return template.JS(fmt.Sprint(result))
}