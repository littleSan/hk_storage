/*
*

	@author:
	@date : 2023/10/19
*/
package pages

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Pagination struct {
	Limit int         `json:"limit" form:"limit" uri:"limit"`
	Page  int         `json:"page" form:"page" uri:"page"`
	Total int         `json:"total" form:"total" uri:"total"`
	Sort  string      `json:"sort" form:"sort" uri:"sort"`
	Data  interface{} `json:"data" form:"data" uri:"data"`
}

func (p *Pagination) InitPage(data interface{}) {
	p.Data = data
}

// Paginate 分页封装 dao层调用
func (p *Pagination) Paginate() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if p.Page == 0 {
			p.Page = 1
		}
		switch {
		case p.Limit > 100:
			p.Limit = 100
		case p.Limit <= 0:
			p.Limit = 10
		}
		offset := (p.Page - 1) * p.Limit
		return db.Offset(offset).Limit(p.Limit)
	}
}

func GeneratePaginationFromRequest(c *gin.Context) (pagination Pagination) {
	if err := c.ShouldBind(&pagination); err != nil {
		fmt.Printf("参数绑定错误:%s\n", err)
	}
	// 校验参数
	if pagination.Limit <= 0 {
		pagination.Limit = 10
	}
	if pagination.Page < 1 {
		pagination.Page = 1
	}

	if len(pagination.Sort) == 0 {
		//pagination.Sort = "timestamp desc"
	}
	return
}
