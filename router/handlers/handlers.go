package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kekaiwang/go-blog/app/admin"
	"github.com/kekaiwang/go-blog/config"
)

// Verify. auth
func Verify(ctx *gin.Context) {
	token := ctx.GetHeader("X-Token")
	if token != config.Get().App.Token {
		// res info
		ctx.JSON(http.StatusOK, admin.Response{
			Data: nil,
			Code: 1,
			Msg:  "please login",
		})
		return
	}

	ctx.Next()
}
