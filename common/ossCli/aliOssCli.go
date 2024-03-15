/*
@author: little
@date : 2023/10/8
*/
package ossCli

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"hk_storage/common/logger"
	"hk_storage/core/configs"
)

var ossClient = &oss.Client{}

func init() {
	fmt.Println("OSS Go SDK Version: ", oss.Version)

	var err error
	ossClient, err = oss.New(configs.TomlConfig.AliOss.Endpoint, configs.TomlConfig.AliOss.Id, configs.TomlConfig.AliOss.Secret)
	if err != nil {
		logger.Info("初始化aliOSS 客户端失败", err)
		return
	}
}
func GetOssCli() *oss.Client {
	return ossClient
}
