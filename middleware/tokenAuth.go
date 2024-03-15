package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"hk_storage/common/constants"
	"hk_storage/common/response"
	"hk_storage/models/sysUser"
	"net/http"
)

func (i *middleware) AuthToken(ctx *gin.Context) (sessionInfo sysUser.SysUser, err error) {
	token := ctx.GetHeader(constants.HeaderLoginToken)
	if token == "" {
		ctx.JSON(http.StatusUnauthorized, response.Failure(ctx, response.AuthorizationError))
		err = errors.New(response.Failure(ctx, response.AuthorizationError).Message)
		return
	}
	if !i.cache.Exits(token) {
		ctx.JSON(http.StatusUnauthorized, response.Failure(ctx, response.AuthorizationError))
		err = errors.New(response.Failure(ctx, response.AuthorizationError).Message)
		return
	}

	cacheData, cacheErr := i.cache.GetLoginInfo(token)
	if cacheErr != nil {
		ctx.JSON(http.StatusUnauthorized, response.FailureMsg(response.AuthorizationError, cacheErr.Error()))
		err = errors.New(cacheErr.Error())
		return
	}

	i.cache.SetLoginInfo(token, cacheData)
	//logger.Info("tokeninfo", cacheData)

	return cacheData, nil

}
