/*
*

	@author:
	@date : 2024/2/28
*/
package studyFile

import "hk_storage/common/pages"

const (
	//待上链
	StatusAll     = 0
	StatusToChain = 1
	//付款成功
	StatusPaySuccess = 2

	//付款失败
	StatusPayFail = 3

	//上链中
	StatusChaining = 4

	//上链成功
	StatusChainSuccess = 5

	//上链失败
	StatusChainFail = 6

	//删除
	StatusDelete = 7
	//付款中
	StatusPaying = 8
)

type StudyFile struct {
	Id           uint64 `json:"id,omitempty" gorm:"column:id" form:"id"`
	Name         string `json:"name,omitempty" gorm:"column:name" form:"name"`
	Uid          string `json:"uid,omitempty" gorm:"column:uid" form:"uid"`
	Url          string `json:"url,omitempty" gorm:"column:url" form:"url"`
	Amount       int    `json:"amount,omitempty" gorm:"column:amount" form:"amount"`                 /**文件数量*/
	Description  string `json:"description',omitempty" gorm:"column:description" form:"description"` /**ERC721*/
	ContractAddr string `json:"contractAddr,omitempty" gorm:"column:contract_addr" form:"contractAddr"`
	Address      string `json:"address,omitempty" gorm:"column:address" form:"address"` //操作者地址
	JsonUrl      string `json:"jsonUrl,omitempty" gorm:"column:json_url" form:"jsonUrl"`
	Token        int64  `json:"token,omitempty" gorm:"column:token" form:"token"`
	Hash         string `json:"hash,omitempty" gorm:"column:hash" form:"hash"`
	PayHash      string `json:"payHash,omitempty" gorm:"column:pay_hash" form:"payHash"` /*付款hash*/
	CreateTime   uint64 `json:"createTime,omitempty" gorm:"column:create_time" form:"createTime"`
	Status       uint   `json:"status,omitempty" gorm:"column:status" form:"status"` /*1 待上链 2 付款成功 3付款失败 4 上链中 5 上链成功 6 上链失败 7  删除*/
}

type Dao interface {
	i()
	TableName() string
	Save(file *StudyFile) (err error)
	Update(file StudyFile) (err error)
	GetFileById(id uint) (study *StudyFile, err error)
	GetFileByHash(hash string) (study *StudyFile, err error)
	List(file StudyFile, page *pages.Pagination) (res []StudyFile, err error)

	GetFileByToken(token int64) (study *StudyFile, err error)
	GetFileByPayHash(hash string) (file *StudyFile, err error)
}

func New() *StudyFile {
	return &StudyFile{}
}

func (d *StudyFile) i() {

}

func (d *StudyFile) TableName() string {
	return "hk_study_file"
}
