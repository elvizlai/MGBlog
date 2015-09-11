package controller
import (
	"gopkg.in/mgo.v2/bson"
	"model/file"
	"enum"
	"fmt"
	"strconv"
	"bytes"
	"image"
	"github.com/nfnt/resize"
	"image/png"
	"image/jpeg"
	"image/draw"
	"io/ioutil"
)

var wm []byte

func init() {
	var err error
	wm, err = ioutil.ReadFile("watermark.png")
	if err != nil {
		panic("watermark.png Not Found")
	}
}

type FileController struct {
	BaseController
	width  int
	height int
	rotate int
}

func (this *FileController) Prepare() {
	this.BaseController.Prepare()

	input := this.Input()
	this.width, _ = strconv.Atoi(input.Get("w"))//宽
	this.height, _ = strconv.Atoi(input.Get("h"))//高
	this.rotate, _ = strconv.Atoi(input.Get("r"))//旋转
}

func (this *FileController) Get() {
	file := file.GetFileById(bson.ObjectIdHex(this.Ctx.Input.Param(":id")))
	if file == nil {
		this.CustomAbort(enum.NotFound.Code(), enum.NotFound.Str())
		return
	}

	//对于请求头中含If-Modified-Since的请求，要对其做出检查
	if m := this.Ctx.Request.Header.Get("If-Modified-Since"); m == fmt.Sprint(file.CreateTime) {
		this.Abort("304")
	}else {
		this.Ctx.ResponseWriter.Header().Add("Cache-Control", "public, max-age=31536000")
		//在资源存在的情况下，为所有的请求返回Last-Modified头
		this.Ctx.ResponseWriter.Header().Add("Last-Modified", fmt.Sprint(file.CreateTime))

		if this.width == 0 && this.height == 0 {
			//添加水印
			this.Ctx.ResponseWriter.Write(waterMark(file.Data))
			return
		}

		img, fileType, _ := image.Decode(bytes.NewBuffer(file.Data))
		m := resize.Resize(uint(this.width), uint(this.height), img, resize.Lanczos3)

		buff := bytes.NewBuffer([]byte{})

		switch fileType {
		case "jpeg":
			jpeg.Encode(buff, m, nil)
		default:
			png.Encode(buff, m)
		}

		this.Ctx.ResponseWriter.Write(waterMark(buff.Bytes()))
	}
}

//水印
func waterMark(picBytes []byte) []byte {
	// 打开水印图并解码
	img, fileType, _ := image.Decode(bytes.NewBuffer(picBytes))

	//读取水印图片
	watermark, _ := png.Decode(bytes.NewBuffer(wm))

	//原始图界限
	origin_size := img.Bounds()

	//创建新图层
	canvas := image.NewNRGBA(origin_size)

	//贴原始图
	draw.Draw(canvas, origin_size, img, image.ZP, draw.Src)
	//贴水印图
	draw.Draw(canvas, watermark.Bounds().Add(image.Pt(origin_size.Dx() - watermark.Bounds().Dx(), origin_size.Dy() - watermark.Bounds().Dy() - 4)), watermark, image.ZP, draw.Over)

	//生成新图片
	buff := bytes.NewBuffer([]byte{})

	switch fileType {
	case "jpeg":
		jpeg.Encode(buff, canvas, &jpeg.Options{95})
	default:
		png.Encode(buff, canvas)
	}

	return buff.Bytes()
}