/*
*

	@author:
	@date : 2024/2/28
*/
package StudyFile

import (
	"hk_storage/common/pages"
	"hk_storage/models/studyFile"
	"time"
)

func (s *service) Save(file *studyFile.StudyFile) (err error) {
	file.CreateTime = uint64(time.Now().UnixMilli())
	return s.StudyFileDao.Save(file)
}

func (s *service) Update(file studyFile.StudyFile) (err error) {
	return s.StudyFileDao.Update(file)
}

func (s *service) GetFileById(id uint) (file *studyFile.StudyFile, err error) {
	return s.StudyFileDao.GetFileById(id)
}
func (s *service) GetFileByHash(hash string) (file *studyFile.StudyFile, err error) {
	return s.StudyFileDao.GetFileByHash(hash)
}
func (s *service) List(file studyFile.StudyFile, page *pages.Pagination) (res []studyFile.StudyFile, err error) {
	return s.StudyFileDao.List(file, page)
}

func (s *service) GetFileByToken(token int64) (study *studyFile.StudyFile, err error) {
	return s.StudyFileDao.GetFileByToken(token)
}

func (s *service) GetFileByPayHash(hash string) (file *studyFile.StudyFile, err error) {
	return s.StudyFileDao.GetFileByPayHash(hash)
}
