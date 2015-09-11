/**
 * Created by elvizlai on 2015/9/6 11:53
 * Copyright © PubCloud
 */
package controller
import (
	"github.com/astaxie/beego"
	"fmt"
)

type Test struct {
	beego.Controller
}

func (this *Test) Get()  {
	fmt.Println(this.Ctx.Input.Params)
	this.Ctx.WriteString(fmt.Sprint(this.Ctx.Input.Params))
}