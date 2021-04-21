package util

import (
	"github.com/jinzhu/gorm"
)

const (
	Page     = "page"  //分頁
	PageSize = "limit" //每頁資料限制筆數
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

	if in, ok := data[PageSize]; ok {
		pageSize = in.(int)
		delete(data, PageSize)
	}
	if in, ok := data[Page]; ok {
		page = in.(int)
		delete(data, Page)
	}
	return data, pageSize, page
}

func TotalSize(in *gorm.DB) int {
	total := 0
	in.Count(&total)
	return total
}
