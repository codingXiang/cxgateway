package util

import (
	"github.com/codingXiang/cxgateway/v3/middleware/pagination"
	"github.com/jinzhu/gorm"
	"strconv"
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

	if in := data[pagination.PageSize]; in != nil {
		pageSize, _ = strconv.Atoi(in.(string))
		delete(data, pagination.PageSize)
	}
	if in := data[pagination.Page]; in != nil {
		page, _ = strconv.Atoi(in.(string))
		delete(data, pagination.Page)
	}
	return data, pageSize, page
}

func TotalSize(in *gorm.DB) int {
	total := 0
	in.Count(&total)
	return total
}
