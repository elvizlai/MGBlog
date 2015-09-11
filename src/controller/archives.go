/**
 * Created by elvizlai on 2015/8/25 16:58
 * Copyright © PubCloud
 */
package controller

type Archives struct {
	BaseController
}

func (this *Archives) Get()  {
	this.TplNames = "archives.html"
}