/*
*

	@author:
	@date : 2023/10/8
*/
package aliOss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"hk_storage/common/ossCli"
	"mime/multipart"
)

var _ Service = (*service)(nil)

type Service interface {
	i()
	Upload(path string, file *multipart.FileHeader) (rest string, err error)
	UploadUrl(path string, url string) (rest string, err error)
}

func New() *service {
	return &service{
		Oss: ossCli.GetOssCli(),
	}
}

type service struct {
	Oss *oss.Client
}

func (s *service) i() {

}
