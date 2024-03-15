package middleware

import (
	"github.com/gin-gonic/gin"
	"hk_storage/common/constants"
	"hk_storage/models/sysUser"
	"net/http"
)

func (i *middleware) WrapAuthHandler(handler func(ctx *gin.Context) (sessionUserInfo sysUser.SysUser, err error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sessionUserInfo, err := handler(ctx)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		ctx.Set(constants.SessionUserInfo, sessionUserInfo)
		return
	}
}
