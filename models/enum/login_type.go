package enum

type LoginType int8

const (
	UserPwdLoginType   LoginType = iota + 1 //
	QQLoginType                             // QQ登录
	EmailLoginType                          // 邮箱登录
	UserPhoneLoginType                      // 手机验证码登录
)
