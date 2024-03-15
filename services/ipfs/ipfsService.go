/*
*

	@author:
	@date : 2023/10/8
*/
package ipfs

import "mime/multipart"

var _ Service = (*service)(nil)

type Service interface {
	i()
	Upload(header *multipart.FileHeader) (rest string, err error)
	UploadJson(obj interface{}, fileName string) (rest string, err error)
}

type service struct {
}

func New() *service {
	return &service{}
}

func (s *service) i() {

}
