/**
 * Author: kekai wang
 */
package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthPing. health ping api
func HealthPing(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "PONG") // pong
}
