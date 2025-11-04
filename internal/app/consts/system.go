package consts

const (
	// RequestIDKey 请求ID在上下文中的键名
	RequestIDKey = "X-Request-ID"
)

const (
	Enable  = 1
	Disable = 2
)

const (
	LoginSystem = 1 // 登录系统
	LoginLogout = 2 // 登出系统

	LoginSuccess = 1 // 登录成功
	LoginFailure = 2 // 登录失败
)
