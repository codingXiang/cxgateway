package util

import (
	"github.com/codingXiang/cxgateway/v3/constants"
	"github.com/jinzhu/gorm"
)



type PaginationData struct {
	Total int         `json:"total"`
	Data  interface{} `json:"data"`
}

func NewPaginationData(total int, data interface{}) *PaginationData {
	return &PaginationData{
		total, data,
	}
}

func Pagination(data map[string]interface{}) func(in *gorm.DB) *gorm.DB {
	return func(in *gorm.DB) *gorm.DB {
		if data == nil {
			return in
		} else {
			var pageSize, page int
			data, pageSize, page = GetLimit(data)
			return in.Limit(pageSize).Offset((page - 1) * pageSize)
		}
	}
}

func GetLimit(data map[string]interface{}) (map[string]interface{}, int, int) {
	var (
		pageSize = 10
		page     = 1
	)

	if in, ok := data[constants.PageSize]; ok {
		pageSize = in.(int)
		delete(data, constants.PageSize)
	}
	if in, ok := data[constants.Page]; ok {
		page = in.(int)
		delete(data, constants.Page)
	}
	return data, pageSize, page
}

func TotalSize(in *gorm.DB) int {
	total := 0
	in.Count(&total)
	return total
}
