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

	if req.Password == "" || req.UserName == "" {
		ApiResponseErr(ctx, errs.ErrInvalidParam)
		return
	}

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
	res.Roles = "admin"

	res.Name = "wkekai"
	res.Avatar = "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"
	res.Introduction = "wkekai blog"
	ApiResponseSuccess(ctx, res)
}
