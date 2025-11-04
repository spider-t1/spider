package verr

// NewErrorSystemLackParam 创建一个系统缺少参数的错误
func NewErrorSystemLackParam(msg string) error {
	return NewError(ErrCodeSystemLackParam, msg)
}

// NewErrorSystemParamConflict 创建一个系统参数冲突的错误
func NewErrorSystemParamConflict(msg string) error {
	return NewError(ErrCodeSystemParamConflict, msg)
}

// NewErrorSystemDataError 创建一个系统数据错误的错误
func NewErrorSystemDataError(msg string) error {
	return NewError(ErrCodeSystemDataError, msg)
}

// NewErrorSystemOperate 创建一个系统操作失败的错误
func NewErrorSystemOperate(msg string) error {
	return NewError(ErrCodeSystemOperate, msg)
}
