package verr

// NewErrorMysqlDataNotExist 创建一个Mysql数据不存在的错误
func NewErrorMysqlDataNotExist(msg string) error {
	return NewError(ErrCodeMysqlDataNotExist, msg)
}

// NewErrorMysqlDataAlreadyExist 创建一个Mysql数据已存在的错误
func NewErrorMysqlDataAlreadyExist(msg string) error {
	return NewError(ErrCodeMysqlDataAlreadyExist, msg)
}

// NewErrorMysqlDataHasChild 创建一个Mysql数据有子级的错误
func NewErrorMysqlDataHasChild(msg string) error {
	return NewError(ErrCodeMysqlDataHasChild, msg)
}
