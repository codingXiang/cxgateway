package middleware

import (
	"fmt"
	"github.com/codingXiang/configer"
	"github.com/codingXiang/go-logger"
	"github.com/codingXiang/gogo-i18n"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"math"
	"os"
	"time"
)

var timeFormat = "2006/05/26:15:04:05 +0800"

// Logger is the logrus logger handler
func Logger() gin.HandlerFunc {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}
	return func(c *gin.Context) {
		// other handler can change c.Path so:
		path := c.Request.URL.Path
		start := time.Now()
		c.Next()
		stop := time.Since(start)
		latency := int(math.Ceil(float64(stop.Nanoseconds()) / 1000000.0))
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		clientUserAgent := c.Request.UserAgent()
		referer := c.Request.Referer()
		dataLength := c.Writer.Size()
		if dataLength < 0 {
			dataLength = 0
		}

		entry := logger.Log.GetLogger().WithFields(logrus.Fields{
			"hostname":   hostname,
			"statusCode": statusCode,
			"latency":    latency, // time to process
			"clientIP":   clientIP,
			"method":     c.Request.Method,
			"path":       path,
			"referer":    referer,
			"dataLength": dataLength,
			"userAgent":  clientUserAgent,
		})

		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		} else {
			msg := fmt.Sprintf("%s - %s [%s] \"%s %s\" %d %d \"%s\" \"%s\" (%dms)", clientIP, hostname, time.Now().Format(timeFormat), c.Request.Method, path, statusCode, dataLength, referer, clientUserAgent, latency)
			if statusCode > 499 {
				entry.Error(msg)
			} else if statusCode > 399 {
				entry.Warn(msg)
			} else {
				entry.Info(msg)
			}
		}
	}
}

func GoI18nMiddleware() gin.HandlerFunc {
	var i18n gogo_i18n.GoGoi18nInterface
	if data, err := configer.Config.GetCore("config").ReadConfig(); err == nil {
		if lang, err := gogo_i18n.LangHandler.GetLanguageTag(data.GetString("i18n.defaultLanguage")); err == nil {
			i18n = gogo_i18n.NewGoGoi18n(lang)
			i18n.SetFileType(data.GetString("i18n.file.type"))
			i18n.LoadTranslationFileArray(data.GetString("i18n.file.path"),
				gogo_i18n.ServerLanguage,
			)
		}
	}
	return func(c *gin.Context) {
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
}
