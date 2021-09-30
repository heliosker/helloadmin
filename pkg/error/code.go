package error

var MessageFlags = map[int]string{
	SUCCESS:             "Ok",
	ERROR:               "啊哦，服务器睡着了",
	ERROR_TOKEN_FAIL:    "Token 不存在或错误",
	ERROR_PASSWORD_FAIL: "用户名或密码错误",
	ERROR_CREATED_FAIL:  "啊哦，创建失败了",
}

func Message(code int) string {
	msg, ok := MessageFlags[code]
	if ok {
		return msg
	}
	return MessageFlags[ERROR]
}
