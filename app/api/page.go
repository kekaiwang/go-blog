package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PageInfo(ctx *gin.Context) {
	ctx.Header("Content-type", "text/html; charset=utf-8")
	ctx.HTML(http.StatusOK, "error.html", nil)
}
