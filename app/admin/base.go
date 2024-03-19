package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kekaiwang/go-blog/internal/common/errs"
)

// Response. res struct
type Response struct {
	Data interface{} `json:"data"` // data
	Msg  string      `json:"msg"`  // msg info
	Code int64       `json:"code"` // res code
	Next interface{} `json:"next"` // next info
}

// ApiResponseSuccess. api res json
func ApiResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Data: data,
		Msg:  "SUCCESS",
		Code: 0,
	})
}

// ApiResponseErr. api res err
func ApiResponseErr(c *gin.Context, err *errs.ErrNo) {
	c.JSON(http.StatusOK, Response{
		Data: nil,
		Msg:  err.Message,
		Code: err.Code,
	})
}

// ApiResponseWarn. api res warn
func ApiResponseWarn(c *gin.Context, err *errs.ErrNo) {
	c.JSON(http.StatusOK, Response{
		Data: nil, // nil data
		Msg:  err.Message,
		Code: err.Code,
		Next: nil,
	})
}

// ApiResponseInfo. api res info
func ApiResponseInfo(c *gin.Context, err *errs.ErrNo) {
	c.JSON(http.StatusOK, Response{
		Data: nil,
		Msg:  err.Message,
		Code: err.Code,
		Next: nil,
	})
}
