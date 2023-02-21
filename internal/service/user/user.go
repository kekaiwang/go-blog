package user

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kekaiwang/go-blog/internal/common/errs"
	"github.com/kekaiwang/go-blog/internal/model"
	"github.com/kekaiwang/go-blog/utils/tools"
)

func (u LoginRequest) Login(ctx *gin.Context) (*LoginResponse, *errs.ErrNo) {
	fmt.Println(ctx.ClientIP())

	// 1. md5 pass
	pass := tools.MD5(u.UserName, u.Password, "$&)w")

	// 2. get user info
	var user model.AdminUser
	info, err := user.GetUser(" name = ? AND password = ? ", []interface{}{u.UserName, pass})
	if err != nil {
		return nil, errs.ErrQueryModel
	}

	res := &LoginResponse{
		ID:       int64(info.ID),
		Password: info.Password,
	}

	// 3. update login info
	info.LastIp = ctx.ClientIP()
	info.LastLogin = time.Now()
	info.LoginCount += 1
	_, err = info.UpdateLogin()
	if err != nil {
		return nil, errs.ErrQueryModel
	}

	return res, nil
}

func (u LoginRequest) AdminLogin(ctx *gin.Context) (*LoginResponse, *errs.ErrNo) {
	fmt.Println(ctx.ClientIP())

	// 1. md5 pass
	pass := tools.MD5(u.UserName, u.Password, "$&)w")

	// 2. get user info
	var user model.AdminUser
	info, err := user.GetUser(" name = ? AND password = ? ", []interface{}{u.UserName, pass})
	if err != nil {
		return nil, errs.ErrQueryModel
	}

	res := &LoginResponse{
		ID:       int64(info.ID),
		Password: info.Password,
	}

	// 3. update login info
	info.LastIp = ctx.ClientIP()
	info.LastLogin = time.Now()
	info.LoginCount += 1
	_, err = info.UpdateLogin()
	if err != nil {
		return nil, errs.ErrQueryModel
	}

	return res, nil
}
