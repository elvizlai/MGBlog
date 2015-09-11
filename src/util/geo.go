/**
 * Created by elvizlai on 2015/8/28 10:01
 * Copyright © PubCloud
 */
package util
import (
	"github.com/astaxie/beego/httplib"
	"fmt"
	"github.com/bitly/go-simplejson"
)

const reqStr = "http://api.map.baidu.com/location/ip?ak=rz6QSfBSl6eNU2dmbTQ1ftGQ&ip=%s&coor=bd09ll"

func InfoGeoByIP(ip string) *simplejson.Json {
	req := httplib.Get(fmt.Sprintf(reqStr, ip))
	resp, err := req.Bytes()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	j, _ := simplejson.NewJson(resp)

	status := j.Get("status").MustInt(1)
	if status == 0 {
		return j
	}else {
		return nil
	}
}