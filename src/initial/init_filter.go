/**
 * Created by elvizlai on 2015/8/25 09:09
 * Copyright © PubCloud
 */
package initial
import (
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego"
	"strings"
	"model/static"
)

var indexFilter = func(ctx *context.Context) {
	static.Add(ctx.Input.IP(), ctx.Input.Url())
}

var filter = func(ctx *context.Context) {
	//过滤静态图片 && 只统计GET请求
	if ctx.Input.Method() == "GET" && !strings.Contains(ctx.Input.Uri(), "/file/") {
		static.Add(ctx.Input.IP(), ctx.Input.Url())
	}
}

func init() {
	beego.InsertFilter("/", beego.BeforeExec, indexFilter)
	beego.InsertFilter("/*", beego.BeforeExec, filter)
}