package util

import (
	"github.com/codingXiang/gogo-i18n"
	"github.com/gin-gonic/gin"
)

func GetI18nData(c *gin.Context) (gogo_i18n.GoGoi18n) {
	data, _ := c.Get("i18n")
	return data.(gogo_i18n.GoGoi18n)
}
