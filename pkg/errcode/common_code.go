package errcode

var (
	Success                     = NewError(0, "成功")
	PasswordFail                = NewError(300000, "账号或密码错误")
	GenerateTokenError          = NewError(300110, "生成 Token 失败")
	UnauthorizedTokenIsNotExist = NewError(300120, "鉴权失败，Token 不存在")
	UnauthorizedTokenError      = NewError(300100, "鉴权失败，Token 错误")
	UnauthorizedTokenTimeOut    = NewError(300200, "鉴权失败，Token 超时")
	TooManyRequests             = NewError(400100, "啊哦，请求过多")
	NotFound                    = NewError(400200, "啊哦，资源不存在")
	InvalidParams               = NewError(400300, "啊哦，请不要瞎填")
	ServerError                 = NewError(500000, "啊哦，服务开小差了")
	CreatedFail                 = NewError(500100, "啊哦，创建失败了")
	SelectedFail                = NewError(500200, "啊哦，查询失败了")
	UpdatedFail                 = NewError(500300, "啊哦，更新失败了")
	DeletedFail                 = NewError(500400, "啊哦，删除失败了")
	UploadFileFail              = NewError(500500, "啊哦，上传文件失败了")
)
