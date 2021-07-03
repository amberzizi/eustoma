package settings

//定义错误码
const (
	Codefail                      = 1000 + iota //"faild"
	CodeSuccess                                 //"success",
	CodeInvalidParam                            //"请求参数有误",
	CodeCheckPasswordWrong                      //密码验证输入错误
	CodeRegisterFail                            //注册失败
	CodeCheckPasswordThroughWrong               //验证密码过程错误
	CodePasswordOrUsernameWrong                 //用户名或密码错误
	CodeVerifyWrong                             //验证码输入错误
	CodeServerBusy                              //服务繁忙
)

var CodeSetting = [...]string{
	1000: "faild",
	1001: "success",
	1002: "请求参数有误",
	1003: "密码验证输入错误",
	1004: "注册失败",
	1005: "验证密码过程错误",
	1006: "用户名或密码错误",
	1007: "验证码输入错误",
	1008: "服务繁忙",
}
