package utility

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func JsonResponse(code int, msg string, data interface{}, c *gin.Context) {
	c.JSON(code, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

func JsonResponseInternalServerError(c *gin.Context) {
	JsonResponse(http.StatusInternalServerError, "Internal server error", nil, c)
}

func JsonSuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "OK",
		"data": data,
	})
}
