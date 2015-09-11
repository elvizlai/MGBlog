package console
import (
	"io/ioutil"
	"model/file"
)

type Upload struct {
	AuthController
}

func (this *Upload) Post() {
	_, header, err := this.GetFile("editormd-image-file")
	if err != nil {
		this.Data["json"] = map[string]interface{}{"success":0, "message":err}
	}else {
		f, _ := header.Open()
		defer f.Close()
		data, _ := ioutil.ReadAll(f)
		pidId := file.AddFile(header.Filename, data, this.CurrentUser)
		this.Data["json"] = map[string]interface{}{"success":1, "message":"OK", "url":"/file/" + pidId.Hex()}
	}
	this.ServeJson()
}