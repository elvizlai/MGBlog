/**
 * Created by elvizlai on 2015/8/25 09:10
 * Copyright © PubCloud
 */
package initial
import (
	"github.com/astaxie/beego"
)

func init() {
	beego.SetLogger("file", `{"filename":"log"}`)

	//beego.SetLevel(beego.LevelInformational)
}