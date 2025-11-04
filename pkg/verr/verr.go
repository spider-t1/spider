package verr

func NewError(code int, msg string) error {
	return &AppError{code: code, msg: msg}
}
func NewErrorWithDetail(code int, msg string, detail any) error {
	return &AppError{code: code, msg: msg, detail: detail}
}

// AppError 在项目中定义统一错误类型
type AppError struct {
	code   int
	msg    string
	detail any
}

func (e *AppError) Code() int {
	return e.code
}

func (e *AppError) Error() string {
	return e.msg
}
func (e *AppError) Detail() any {
	return e.detail
}

func (e *AppError) ErrorWithCodeStr() string {
	return ErrCodeMap[e.code] + " : " + e.msg
}
