package middleware

import (
	"github.com/gin-gonic/gin"
	"hk_storage/common/redisCli/userRedis"
	"hk_storage/models/sysUser"
)

var _ Middleware = (*middleware)(nil)

type Middleware interface {
	i()
	Cors() gin.HandlerFunc
	AuthToken(ctx *gin.Context) (sessionUserInfo sysUser.SysUser, err error)
	WrapAuthHandler(handler func(ctx *gin.Context) (sessionUserInfo sysUser.SysUser, err error)) gin.HandlerFunc
}

func New() *middleware {
	return &middleware{
		cache: userRedis.New(),
	}
}

type middleware struct {
	cache userRedis.Service
}

func (i *middleware) i() {

}
