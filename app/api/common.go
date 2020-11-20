/**
 * Author: kekai wang
 */
package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthPing(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "PONG")
	return
}
