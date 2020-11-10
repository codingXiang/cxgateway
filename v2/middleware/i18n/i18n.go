package i18n

import (
	"github.com/codingXiang/configer/v2"
	"github.com/codingXiang/cxgateway/v2/middleware"
	"github.com/codingXiang/cxgateway/v2/server"
	gogo_i18n "github.com/codingXiang/gogo-i18n"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type I18n struct {
	config *viper.Viper
}

func New(config *viper.Viper) middleware.Object {
	return &I18n{
		config: config,
	}
}

//Version
func (r *I18n) Handle() gin.HandlerFunc {
	if r.config.GetBool(configer.GetConfigPath(server.I18n, server.Enable)) {
		var i18n gogo_i18n.GoGoi18nInterface
		if gogo_i18n.LangHandler == nil {
			gogo_i18n.LangHandler = gogo_i18n.NewLanguageHandler()
		}
		if lang, err := gogo_i18n.LangHandler.GetLanguageTag(r.config.GetString(configer.GetConfigPath(server.I18n, server.DefaultLanguage))); err == nil {
			i18n = gogo_i18n.NewGoGoi18n(lang)
			i18n.SetFileType(r.config.GetString("i18n.file.type"))
			i18n.LoadTranslationFileArray(r.config.GetString("i18n.file.path"),
				gogo_i18n.ServerLanguage,
			)
		} else {
			panic(err)
		}
		return func(c *gin.Context) {
			if i18n == nil {
				panic("i18n is not set")
				c.Abort()
			}

			locale := c.Query("locale")
			if locale != "" {
				c.Request.Header.Set("Accept-Language", locale)
			}
			if lang, err := gogo_i18n.LangHandler.GetLanguageTag(c.GetHeader("Accept-Language")); err == nil {
				i18n.SetUseLanguage(lang)
				c.Set("i18n", i18n)
				c.Next()
			} else {
				c.Abort()
			}
		}
	} else {
		return func(c *gin.Context) {
			c.Next()
		}
	}

}
