package util

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func IsViolatingUniqueConstraint(err error) bool {
	return strings.Contains(strings.ToLower(err.Error()), "unique")
}

func ParamErrResponse(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"code": http.StatusBadRequest,
		"msg":  "参数错误",
		"data": nil,
	})
}

func InternalErrResponse(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"code": http.StatusInternalServerError,
		"msg":  "服务器内部错误",
		"data": nil,
	})
}
