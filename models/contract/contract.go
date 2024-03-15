/*
*

	@author:
	@date : 2023/11/9
*/
package contract

import (
	"hk_storage/core/database"
)

type Contract struct {
	Id         uint64 `json:"id,omitempty" gorm:"column:id" form:"id"`
	From       string `json:"from,omitempty" gorm:"column:from" form:"from"`
	PriKey     string `json:"priKey,omitempty" gorm:"column:pri_key" form:"priKey"`
	Address    string `json:"address,omitempty" gorm:"column:address" form:"address"`
	Type       string `json:"type',omitempty" gorm:"column:type" form:"type"` /**ERC721*/
	Abi        string `json:"abi,omitempty" gorm:"column:abi" form:"abi"`
	Code       string `json:"code,omitempty" gorm:"column:code" form:"code"`
	Mark       string `json:"mark,omitempty" gorm:"column:mark" form:"mark"`
	CreateTime uint64 `json:"createTime,omitempty" gorm:"column:create_time" form:"createTime"`
}

type ABIQueryDTO struct {
	ContractAddr string        `json:"contractAddr" form:"contractAddr"`
	PriKey       string        `json:"priKey" form:"priKey"`
	From         string        `json:"from" form:"from"`
	AbiCode      string        `json:"abiCode" form:"abiCode"`
	Method       string        `json:"method" form:"method"`
	Params       []interface{} `json:"params" form:"params"`
	Amount       string        `json:"amount" form:"amount"`
}

type ABIResponse struct {
	Code    int                    `json:"code" form:"code"`
	Message string                 `json:"message" form:"message"`
	Data    map[string]interface{} `json:"data" form:"data"`
}

func (d *Contract) TableName() string {
	return "hk_contract"
}

func (d *Contract) List() (res []Contract, err error) {
	err = database.GetDB().Model(d).Where(d).Find(d).Error
	return
}

func (d *Contract) GetInfo() (err error) {
	err = database.GetDB().Model(d).Where(d).First(d).Error
	return
}
