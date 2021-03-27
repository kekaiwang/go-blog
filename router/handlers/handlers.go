package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kekaiwang/go-blog/app/admin"
	"github.com/kekaiwang/go-blog/config"
)

func Verify(ctx *gin.Context) {
	token := ctx.GetHeader("X-Token")
	if token != config.Get().App.Token {
		ctx.JSON(http.StatusOK, admin.Response{
			Data: nil,
			Code: 1,
			Msg:  "请先登陆",
		})
		return
	}

	ctx.Next()
}
