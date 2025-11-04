package response

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"spider/internal/app/consts"
	"spider/internal/app/types/types_common"
	"spider/pkg/verr"
)

type buildConfig struct {
	data       interface{}
	sourceData interface{}
	code       int32
	msg        string
	httpCode   int
	sendErrMsg bool
}

type BuildOption func(*buildConfig)

func WithData(data any) BuildOption {
	return func(c *buildConfig) {
		c.data = data
	}
}

func WithListData(req types_common.IBaseListResp) BuildOption {
	return func(c *buildConfig) {
		if req != nil && !reflect.ValueOf(req).IsNil() {
			req.Adjust()
		}
		c.data = req
	}
}

func WithErrCode(code int32) BuildOption {
	return func(c *buildConfig) {
		c.code = code
	}
}

func WithSourceData(data interface{}) BuildOption {
	return func(c *buildConfig) {
		c.sourceData = data
	}
}

func WithHTTPCode(code int) BuildOption {
	return func(c *buildConfig) {
		c.httpCode = code
	}
}

func WithSendErrMsg() BuildOption {
	return func(c *buildConfig) {
		c.sendErrMsg = true
	}
}

type Response struct {
	Code      int         `json:"code"`           // 业务码
	Msg       string      `json:"msg"`            // 提示信息
	RequestId string      `json:"requestId"`      // 请求ID
	Data      interface{} `json:"data,omitempty"` // 数据内容（成功时）
}

type qwError struct {
	TimeStamp int64       `json:"timeStamp"`
	Code      int         `json:"code"`
	Api       string      `json:"api"`
	Msg       string      `json:"msg"`
	Stack     string      `json:"stack"`
	Request   interface{} `json:"request"`
}

// Success 成功返回
func Success(ctx *gin.Context, data interface{}) {
	rid, _ := ctx.Get(consts.RequestIDKey)
	ctx.JSON(http.StatusOK, Response{
		Code:      http.StatusOK,
		Msg:       "success",
		Data:      data,
		RequestId: rid.(string),
	})
}

// HandleDefault ，返回延迟处理函数
func HandleDefault(ctx *gin.Context, opts ...BuildOption) func(*error, any) {
	// 定义延迟处理函数
	handler := func(err *error, r any) {

		rid, _ := ctx.Get(consts.RequestIDKey)

		conf := &buildConfig{}

		if opts != nil {
			for _, opt := range opts {
				opt(conf)
			}
		}
		if r != nil {
			*err = verr.NewError(verr.ErrCodeInternal, fmt.Sprintf("%v", r))
		}
		if *err != nil {
			errVal := fmt.Sprintf("%+v", *err)
			code := verr.ErrCodeInternal
			var e *verr.AppError
			if errors.As(*err, &e) {
				code = e.Code()
				if e.Error() != "" {
					errVal = e.ErrorWithCodeStr()
				}
			}
			ctx.JSON(http.StatusOK, Response{
				Code:      code,
				Msg:       errVal,
				RequestId: rid.(string),
			})
			return
		}
		if conf.sourceData != nil {
			Success(ctx, conf.sourceData)
			return
		}
		Success(ctx, conf.data)
	}

	return handler
}

//func HandleListDefault(ctx *gin.Context, res common.IBaseListResp) func(*error, any) {
//	// 定义延迟处理函数
//	handler := func(err *error, r any) {
//		if r != nil {
//			*err = errors.New(fmt.Sprintf("%v", r))
//		}
//		if *err != nil {
//			resValue := fmt.Sprintf("%v", res)
//			code := http.StatusInternalServerError
//			var e *AppError
//			if errors.As(*err, &e) {
//				code = e.Code()
//				if e.Error() != "" {
//					resValue = e.Error()
//				}
//			}
//			Error(ctx, err, code, resValue)
//			return
//		}
//		res.Adjust()
//		Success(ctx, res)
//	}
//
//	return handler
//}
