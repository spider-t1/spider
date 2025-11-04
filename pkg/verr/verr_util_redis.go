package verr

// NewErrorCodeRedisDataNotExist 创建一个Redis数据不存在的错误
func NewErrorCodeRedisDataNotExist(msg string) error {
	return NewError(ErrCodeRedisDataNotExist, msg)
}

// NewErrorCodeRedisDataAlreadyExist 创建一个Redis数据已存在的错误
func NewErrorCodeRedisDataAlreadyExist(msg string) error {
	return NewError(ErrCodeRedisDataAlreadyExist, msg)
}

// NewErrorDataLockNotHold 创建一个数据锁未持有错误
func NewErrorDataLockNotHold(msg string) error {
	return NewError(ErrCodeRedisLockNotHold, msg)
}
