package error

var MessageFlags = map[int]string{
	SUCCESS:             "Ok",
	ERROR:               "Fail",
	ERROR_TOKEN_FAIL:    "Token 不存在或错误",
	ERROR_PASSWORD_FAIL: "用户名或密码错误",
}

func Message(code int) string {
	msg, ok := MessageFlags[code]
	if ok {
		return msg
	}
	return MessageFlags[ERROR]
}
