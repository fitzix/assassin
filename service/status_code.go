package service

// A-B-CC
const (
	StatusSuccess      = 0
	StatusBadRequest   = 1000
	StatusParamErr     = 1001
	StatusUserNotExist = 2000
	StatusUserWrongPwd = 2001
)

var statusText = map[int]string{
	StatusSuccess:      "ok",
	StatusBadRequest:   "请求出错",
	StatusParamErr:     "参数错误",
	StatusUserWrongPwd: "用户名或密码错误",
	StatusUserNotExist: "用户不存在", // 用户
}

// StatusText returns a text for the HTTP status code. It returns the empty
// string if the code is unknown.
func AsnStatusText(code int) string {
	return statusText[code]
}
