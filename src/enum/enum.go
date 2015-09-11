package enum
import "fmt"

type Code int

const (
	OK Code = iota
)

//100~ 用户相关
const (
	EmailAlreadyExist Code = iota + 100
	NickNameAlreadyExist
	PasswordIncorrect
	UserNotExist
)

const (
	BadRequest Code = 400
	UnAuthorized Code = 401
	NotFound Code = 404
	UNKNOWN Code = 999
)

func (c Code)Str() string {
	switch c{
	case OK:
		return "OK"
	case NickNameAlreadyExist:
		return "用户名已存在"
	case EmailAlreadyExist:
		return "邮箱已存在"
	case PasswordIncorrect:
		return "密码错误"
	case UserNotExist:
		return "用户不存在"
	case UnAuthorized:
		return "未授权"
	case BadRequest:
		return "Bad Request"
	case NotFound:
		return "Not Found"
	default:
		return "UNKNOWN:" + fmt.Sprint(c)
	}
}

func (c Code)Code() int {
	return int(c)
}