package errcode

var (
	Success                   = NewError(200, "成功")
	ServerError               = NewError(500, "服务内部错误")
	ErrorUserNameUsed         = NewError(10000001, "用户名已存在")
	ErrorUserInformationWrong = NewError(10000002, "用户名或者密码错误")
	ErrorUserNotExist         = NewError(10000003, "用户不存在")
	ErrorTokenExist           = NewError(10000004, "Token不存在")
	ErrorTokenTimeout         = NewError(10000005, "Token超时")
	ErrorTokenTypeWrong       = NewError(10000006, "Token格式错误")
)
