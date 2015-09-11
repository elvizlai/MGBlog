package api

import (
	"github.com/astaxie/beego/validation"
	"enum"
	"model/user"
	"strings"
	"github.com/astaxie/beego"
	"util"
	"controller"
)

type User struct {
	controller.BaseController
}

//用户注册
func (this *User) Register() {
	req := this.ReqJson()
	if req != nil {
		email := req.Get("email").MustString()
		nickName := req.Get("nickName").MustString()
		password := req.Get("password").MustString()

		valid := validation.Validation{}
		valid.Email(email, "email")
		valid.MinSize(nickName, 6, "nickNameMin")
		valid.MaxSize(nickName, 12, "nickNameMax")
		valid.MinSize(password, 6, "passwordMin")
		valid.MaxSize(password, 12, "passwordMax")

		if valid.HasErrors() {
			this.CustomAbort(enum.BadRequest.Code(), enum.BadRequest.Str())
		}

		err := user.AddUser(email, nickName, password)
		if err == nil {
			this.RespJson(enum.OK, nil)
		}else {
			if strings.Contains(err.Error(), "email") {
				this.RespJson(enum.EmailAlreadyExist, nil)
			}else if strings.Contains(err.Error(), "nickname") {
				this.RespJson(enum.NickNameAlreadyExist, nil)
			}else {
				beego.Error(err)
			}
		}
	}
}

//修改密码
//func (this *User) ChangePwd() {
//	req := this.ReqJson()
//	if req != nil {
//		if this.CruUser == nil {
//			this.RespJson(enum.UnAuthorized, nil)
//		}else {
//			oriPwd := req.Get("oriPwd").MustString()
//			newPwd := req.Get("newPwd").MustString()
//
//			valid := validation.Validation{}
//
//			valid.MinSize(oriPwd, 6, "oriPwdMin")
//			valid.MaxSize(oriPwd, 12, "oriPwdMax")
//			valid.MinSize(newPwd, 6, "newPwdMin")
//			valid.MaxSize(newPwd, 12, "newPwdMax")
//
//			if valid.HasErrors() {
//				this.CustomAbort(enum.BadRequest.Code(), enum.BadRequest.Str())
//			}
//
//			if util.Md5(this.CruUser.Salt + oriPwd) != this.CruUser.Password {
//				this.RespJson(enum.PasswordIncorrect, nil)
//			}else {
//				user.ChangePwd(this.CruUser.Email, oriPwd, newPwd)
//				this.RespJson(enum.OK, nil)
//			}
//		}
//	}
//}

//修改昵称


//用户登录
func (this *User) Login() {
	json := this.ReqJson()
	if json != nil {
		email := json.Get("email").MustString()
		password := json.Get("password").MustString()

		valid := validation.Validation{}
		valid.Email(email, "email")
		valid.MinSize(password, 6, "passwordMin")
		valid.MaxSize(password, 12, "passwordMax")

		if valid.HasErrors() {
			this.CustomAbort(enum.BadRequest.Code(), enum.BadRequest.Str())
		}

		u := user.GetUserByEmail(email)
		if u == nil {
			//用户不存在
			this.RespJson(enum.UserNotExist, nil)
		}else if util.Md5(u.Salt + password) != u.Password {
			//密码错误
			this.RespJson(enum.PasswordIncorrect, nil)
		}else {
			this.SetSession("uId", u.Id.Hex())
			user.SetToken(email, this.StartSession().SessionID())//将浏览器的cookie token写入库中，做异地登录验证
			this.RespJson(enum.OK, map[string]interface{}{"url":"/"})
		}
	}
}