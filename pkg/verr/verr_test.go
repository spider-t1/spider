package verr

import (
	"errors"
	"fmt"
	"net/http"
	"testing"
)

func TestJErr(t *testing.T) {
	err := NewError(http.StatusNotFound, "没有权限")
	adjust(&err)
}

func adjust(err *error) {
	var e *AppError
	if errors.As(*err, &e) { // 使用指针进行类型断言
		// 如果转换成功，可以使用 appErr
		fmt.Println("Error code:", e.Code())
		fmt.Println("Error message:", e.Error())
	}
}
