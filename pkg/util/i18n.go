package util

import (
	"github.com/codingXiang/go-logger"
	"github.com/codingXiang/gogo-i18n"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/text/language"
)

func GetI18nData(c *gin.Context) (gogo_i18n.GoGoi18nInterface) {
	data, _ := c.Get("i18n")
	return data.(*gogo_i18n.GoGoi18n)
}

func LoadI18nDataFromDatabase(db *gorm.DB, storeType string, storePath string) {
	//設定多語系
	i18ns := make([]*gogo_i18n.GoGoi18nMessage, 0)
	i18nMsgs := make([]gogo_i18n.GoGoi18nMessageInterface, 0)
	logger.Log.Info("load i18n data from database...")
	if err := db.Find(&i18ns).Error; err == nil {
		for _, i18n := range i18ns {
			msg := createI18nMessageFromRecord(i18n)
			i18nMsgs = append(i18nMsgs, msg)
		}
		logger.Log.Info("save i18n data to", storePath, "type = ", storeType)
		gogo_i18n.StoreDataToFile(storeType, storePath, i18nMsgs)
	} else {
		logger.Log.Error("load i18n data from database failed, err =", err.Error())
	}
}

func createI18nMessageFromRecord(i18n *gogo_i18n.GoGoi18nMessage) gogo_i18n.GoGoi18nMessageInterface {
	logger.Log.Debug("load i18n record =", i18n)
	switch i18n.Language {
	case "zh-Hant":
		return gogo_i18n.NewGoGoi18nMessage(language.TraditionalChinese, i18n.Key, i18n.Value)
	case "zh-Hans":
		return gogo_i18n.NewGoGoi18nMessage(language.SimplifiedChinese, i18n.Key, i18n.Value)
	case "es":
		return gogo_i18n.NewGoGoi18nMessage(language.English, i18n.Key, i18n.Value)
	default:
		return gogo_i18n.NewGoGoi18nMessage(language.TraditionalChinese, i18n.Key, i18n.Value)
	}
}
