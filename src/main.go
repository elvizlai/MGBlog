package main

import (
	_ "initial"
	_ "router"
	"github.com/astaxie/beego"
)

func main() {
	beego.Debug(beego.VERSION)
	beego.Run()
}