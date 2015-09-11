package console
import (
	"model/article"
	"strings"
	"enum"
)

type NewArticle struct {
	AuthController
}

func (this *NewArticle) Get() {
	this.TplNames = "console/new_article.html"
}

func (this *NewArticle) Post() {
	req := this.ReqJson()

	title := req.Get("title").MustString()
	tagStr := req.Get("tags").MustString()

	var tags []string
	if tagStr == "" {
		tags = []string{"未分类"}
	}else if strings.Contains(tagStr, ";") {
		tags = strings.Split(tagStr, ";")
	}else {
		tags = strings.Split(tagStr, "；")
	}

	for i, _ := range tags {
		tags[i] = strings.TrimLeft(tags[i], " ")
		tags[i] = strings.TrimRight(tags[i], " ")
	}

	markdown := req.Get("markdown").MustString()
	htmlContent := req.Get("htmlContent").MustString()

	err := article.AddArticle(this.CurrentUser.Id, title, tags, markdown, htmlContent)

	if err == nil {
		this.RespJson(enum.OK, nil)
	}else {
		this.RespJson(enum.UNKNOWN, err)
	}
}