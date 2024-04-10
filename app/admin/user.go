package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/kekaiwang/go-blog/internal/common/errs"
	"github.com/kekaiwang/go-blog/internal/service/user"
)

// Login. admin login
func Login(ctx *gin.Context) {
	var req user.LoginRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ApiResponseErr(ctx, errs.ErrBindJson)
		return
	}

	// check params & return res
	if req.Password == "" || req.UserName == "" {
		ApiResponseErr(ctx, errs.ErrInvalidParam) // res
		return
	}

	// login admin
	res, err := req.Login(ctx)
	if err != nil {
		ApiResponseErr(ctx, err)
		return
	}

	ApiResponseSuccess(ctx, res)
}

// Info base info
func Info(ctx *gin.Context) {
	var res user.InfoResponse
	res.Roles = "admin" // role

	res.Name = "wkekai"
	res.Avatar = "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif" // avatar link
	res.Introduction = "wkekai blog"                                                   // introduce
	ApiResponseSuccess(ctx, res)
}
