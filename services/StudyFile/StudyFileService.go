/*
*

	@author:
	@date : 2024/2/28
*/
package StudyFile

import (
	"hk_storage/common/pages"
	"hk_storage/models/studyFile"
)

var _ Service = (*service)(nil)

type Service interface {
	i()
	Save(file *studyFile.StudyFile) (err error)
	Update(file studyFile.StudyFile) (err error)
	List(file studyFile.StudyFile, page *pages.Pagination) (res []studyFile.StudyFile, err error)
	GetFileById(id uint) (study *studyFile.StudyFile, err error)
	GetFileByHash(hash string) (study *studyFile.StudyFile, err error)
	GetFileByToken(token int64) (study *studyFile.StudyFile, err error)
	GetFileByPayHash(hash string) (file *studyFile.StudyFile, err error)
}

type service struct {
	StudyFileDao studyFile.Dao
}

func New() *service {
	return &service{
		StudyFileDao: studyFile.New(),
	}
}

func (s *service) i() {
}
