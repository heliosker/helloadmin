package error

var MessageFlags = map[int]string{
	SUCCESS:          "Ok",
	ERROR:            "Fail",
	ERROR_TOKEN_FAIL: "Token 不存在或错误",
}

func GetMessage(code int) string {
	msg, ok := MessageFlags[code]
	if ok {
		return msg
	}
	return MessageFlags[ERROR]
}
