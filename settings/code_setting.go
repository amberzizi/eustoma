package settings

//定义错误码
const (
	Codefail                               = 1000 + iota //"faild"
	CodeSuccess                                          //"success",
	CodeInvalidParam                                     //"请求参数有误",
	CodeCheckPasswordWrong                               //密码验证输入错误
	CodeRegisterFail                                     //注册失败
	CodeCheckPasswordThroughWrong                        //验证密码过程错误
	CodePasswordOrUsernameWrong                          //用户名或密码错误
	CodeVerifyWrong                                      //验证码输入错误
	CodeUserExist                                        //用户已存在
	CodeUserNotExist                                     //用户不存在
	CodeServerBusy                                       //服务繁忙
	ErrorInvalidToken                                    //无效token
	ErrorGenToken                                        //生成token失败
	ErrorEmptyToken                                      //无认证token
	ErrorFormatToken                                     //token格式错误
	ErrorUserNotLogin                                    //用户未登录
	ErrorAccessTokenExpiredOutOfTime                     //accesstoken过期
	ErrorAccessTokenSingleLoginCheck                     //单点登录检查异常-查询异常
	ErrorAccessTokenSingleLoginHavefreshed               //单点登录检查异常-已有新的登录token-请重新登录
	CodeOutOfRange                                       //超出限制
	CodePostError                                        //发布失败
	CodeVoteError                                        //投票失败
	ErrorVoteOutOfTime                                   //超出投票时间
	ErrorVoteRepeat                                      //重复投票
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
	1008: "用户已存在",
	1009: "用户不存在",
	1010: "服务繁忙",
	1011: "无效token",
	1012: "生成token失败",
	1013: "无认证token,请登录",
	1014: "token格式错误",
	1015: "用户未登录",
	1016: "accesstoken过期",
	1017: "单点登录检查异常-查询异常",
	1018: "单点登录检查异常-已有新的登录token-请重新登录",
	1019: "超出限制",
	1020: "发布失败",
	1021: "投票失败",
	1022: "超出投票时间",
	1023: "请勿重复投票",
}
