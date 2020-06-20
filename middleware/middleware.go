package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/codingXiang/cxgateway/pkg/e"
	"github.com/codingXiang/go-logger"
	"github.com/codingXiang/gogo-i18n"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"strconv"
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

func GoI18nMiddleware(data *viper.Viper) gin.HandlerFunc {
	var i18n gogo_i18n.GoGoi18nInterface
	if lang, err := gogo_i18n.LangHandler.GetLanguageTag(data.GetString("i18n.defaultLanguage")); err == nil {
		i18n = gogo_i18n.NewGoGoi18n(lang)
		i18n.SetFileType(data.GetString("i18n.file.type"))
		i18n.LoadTranslationFileArray(data.GetString("i18n.file.path"),
			gogo_i18n.ServerLanguage,
		)
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
}

func GoCors(data *viper.Viper) gin.HandlerFunc {
	var (
		allowOrigins = data.GetStringSlice("cors.allowOrigins")
		allowMethods = data.GetStringSlice("cors.allowMethods")
		allowHeaders = data.GetStringSlice("cors.allowHeaders")
	)
	logger.Log.Info("go cors")
	config := cors.DefaultConfig()
	config.AllowOrigins = allowOrigins
	config.AllowHeaders = allowHeaders
	config.AllowMethods = allowMethods
	return cors.New(config)
}

//GoCache 是否允許存取快取（realtime = true 時存取快取）
func GoCache(c *gin.Context) {
	var (
		enableCache = true
	)
	if tmp, isExist := c.GetQuery("realtime"); isExist {
		if cache, err := strconv.ParseBool(tmp); err == nil {
			enableCache = !cache
		}
	}
	c.Set("enableCache", enableCache)
	c.Next()
}

//PermissionAuth 權限驗證
func PermissionAuth(config, auth *viper.Viper) gin.HandlerFunc {
	var (
		//取得 auth server 資料
		permissionApi = auth.GetString("auth.server") + auth.GetString("auth.permissionCheck.path")
		method        = auth.GetString("auth.permissionCheck.method")
		// 取得 application 資料
		targetApp = config.GetString("application.appId")
	)
	return func(c *gin.Context) {
		client := &http.Client{}

		//對 permission 的 api 進行存取
		req, err := http.NewRequest(method, permissionApi, nil)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, e.UnknownError(err.Error()))
			return
		}

		//設定權限控制相關 header
		req.Header.Set("auth-app", c.GetHeader("auth-app"))
		req.Header.Set("auth-token", c.GetHeader("auth-token"))
		req.Header.Set("target-app", targetApp)
		req.Header.Set("auth-path", c.Request.URL.Path)
		req.Header.Set("auth-method", c.Request.Method)

		//送出 request
		resp, err := client.Do(req)

		if resp == nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"errMsg": "Connect to authority service failed.",
			})
			return
		}

		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			c.Next()
			return
		}

		//讀取 response body
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}
		//將 response body 轉換成 map
		var response = new(map[string]interface{})
		if err := json.Unmarshal(bodyBytes, response); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
}
