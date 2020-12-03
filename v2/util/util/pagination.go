package util

import (
	"github.com/codingXiang/cxgateway/v2/middleware/pagination"
	"github.com/jinzhu/gorm"
)

func Pagination(in *gorm.DB, data map[string]interface{}) *gorm.DB {
	if data == nil {
		return in
	} else if data[pagination.PAGE_SIZE] == nil && data[pagination.PAGE] == nil {
		return in
	} else {
		var (
			pageSize = 10
			page     = 1
		)

		if in := data[pagination.PAGE_SIZE]; in != nil {
			pageSize = in.(int)
		}
		if in := data[pagination.PAGE_SIZE]; in != nil {
			page = in.(int)
		}
		return in.Limit(pageSize).Offset((page - 1) * pageSize)
	}
}
