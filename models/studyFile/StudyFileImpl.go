/*
*

	@author:
	@date : 2024/2/28
*/
package studyFile

import (
	"fmt"
	"hk_storage/common/logger"
	"hk_storage/common/pages"
	"hk_storage/core/database"
	"strings"
)

func (d *StudyFile) Save(file *StudyFile) (err error) {
	err = database.GetDB().Model(d).Save(file).Error
	if err != nil {
		logger.Info("保存数据失败", err)
	}
	return err
}

func (d *StudyFile) Update(file StudyFile) (err error) {
	err = database.GetDB().Model(d).Omit("id", "uid").Where("id= ?", file.Id).Updates(&file).Error
	logger.Info("更新hash信息", err)
	return err
}

func (d *StudyFile) GetFileById(id uint) (file *StudyFile, err error) {
	err = database.GetDB().Model(d).Where("id = ?", id).First(&file).Error
	if err != nil {
		return nil, err
	}
	return file, err
}

func (d *StudyFile) GetFileByHash(hash string) (file *StudyFile, err error) {
	err = database.GetDB().Model(d).Where("hash = ?", hash).First(&file).Error
	if err != nil {
		return nil, err
	}
	return file, err
}

func (d *StudyFile) List(file StudyFile, page *pages.Pagination) (res []StudyFile, err error) {
	db := database.GetDB().Model(d)
	if strings.Trim(file.Uid, " ") != "" {
		db.Where("uid = ?", file.Uid)
	}
	if strings.Trim(file.Address, " ") != "" {
		db.Where("address = ?", file.Address)
	}
	if strings.Trim(file.Hash, " ") != "" {
		db.Where("hash = ?", file.Hash)
	}
	if strings.Trim(file.PayHash, " ") != "" {
		db.Where("pay_hash = ?", file.PayHash)
	}

	if file.Status > 0 {
		db.Where("status = ? ", file.Status)
	} else {
		db.Where("status != ?", StatusDelete)
	}
	if file.Token > 0 {
		db.Where("token = ? ", file.Token)
	}
	if strings.Trim(file.Name, " ") != "" {
		db.Where("name like?", fmt.Sprintf("%%%v%%", file.Name))
	}

	var total int64
	err = db.Count(&total).Error
	page.Total = int(total)
	err = db.Scopes(page.Paginate()).Order(page.Sort).Find(&res).Error
	return
}

func (d *StudyFile) GetFileByToken(token int64) (study *StudyFile, err error) {
	err = database.GetDB().Model(d).Where("token = ?", token).First(&study).Error
	if err != nil {
		return nil, err
	}
	return study, err
}

func (d *StudyFile) GetFileByPayHash(hash string) (file *StudyFile, err error) {
	err = database.GetDB().Model(d).Where("pay_hash = ?", hash).First(&file).Error
	if err != nil {
		return nil, err
	}
	return file, err
}
