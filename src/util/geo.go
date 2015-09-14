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

const reqStr = "http://ip-api.com/json/%s"

func InfoGeoByIP(ip string) *simplejson.Json {
	req := httplib.Get(fmt.Sprintf(reqStr, ip))
	resp, err := req.Bytes()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	j, _ := simplejson.NewJson(resp)

	status := j.Get("status").MustString()
	if status == "success" {
		return j
	}else {
		return nil
	}
}