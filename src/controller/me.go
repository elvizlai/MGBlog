/**
 * Created by elvizlai on 2015/7/28 17:25
 * Copyright © PubCloud
 */
package controller

type Me struct {
	BaseController
}

//关于我
func (this *Me) Get() {
	this.TplNames = "me.html"
}
