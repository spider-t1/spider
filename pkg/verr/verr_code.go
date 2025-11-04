package verr

// HTTP 状态码
const (
	ErrCodeUnauthorized    = 401 // 缺少认证
	ErrCodeForbidden       = 403 // 无操作权限
	ErrCodeNotFound        = 404 // 路径不存在
	ErrCodeTooManyRequests = 429 // 访问频繁

	ErrCodeInternal = 500 // 服务器开小差了
)

// 业务方面错误码
//
//	ErrCodeSystem 400001-400099 业务级错误
const (
	ErrCodeSystemLackParam     = 400001 // 缺少参数
	ErrCodeSystemParamConflict = 400002 // 参数存在冲突
	ErrCodeSystemDataError     = 400003 // 数据错误
	ErrCodeSystemOperate       = 400004 // 操作错误
)

// ErrCodeMysql 400101-400199 Mysql错误
const (
	ErrCodeMysqlDataNotExist     = 400101 // mdb数据不存在
	ErrCodeMysqlDataAlreadyExist = 400102 // mdb数据已存在
	ErrCodeMysqlDataHasChild     = 400103 // mdb数据存在子级
)

// ErrCodeRedis 400201-400299 Redis错误
const (
	ErrCodeRedisDataNotExist     = 400201 // rdb数据不存在
	ErrCodeRedisDataAlreadyExist = 400202 // rdb数据已存在
	ErrCodeRedisLockNotHold      = 400203 // rdb数据锁未持有
)

var ErrCodeMap = map[int]string{
	// HTTP 状态码
	ErrCodeUnauthorized:    "未授权",
	ErrCodeForbidden:       "无权限",
	ErrCodeNotFound:        "路径不存在",
	ErrCodeTooManyRequests: "访问频繁",
	ErrCodeInternal:        "服务器开小差了",

	//	ErrCodeSystem 400001-400099 业务级错误
	ErrCodeSystemLackParam:     "缺少参数",
	ErrCodeSystemParamConflict: "参数存在冲突",
	ErrCodeSystemDataError:     "数据错误",
	ErrCodeSystemOperate:       "操作错误",

	// ErrCodeMysql 400101-400199 Mysql错误
	ErrCodeMysqlDataNotExist:     "mdb数据不存在",
	ErrCodeMysqlDataAlreadyExist: "mdb数据已存在",
	ErrCodeMysqlDataHasChild:     "mdb数据存在子级",

	// ErrCodeRedis 400201-400299 Redis错误
	ErrCodeRedisDataNotExist:     "rdb数据不存在",
	ErrCodeRedisDataAlreadyExist: "rdb数据已存在",
	ErrCodeRedisLockNotHold:      "rdb数据锁未持有",
}
