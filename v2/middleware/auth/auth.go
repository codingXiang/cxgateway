package auth

import (
	"encoding/json"
	"github.com/codingXiang/configer/v2"
	"github.com/codingXiang/cxgateway/v2/middleware"
	"github.com/codingXiang/cxgateway/v2/model"
	"github.com/codingXiang/cxgateway/v2/util/e"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	_symbol          = "-"
	_auth            = "auth"
	_enable          = "enable"
	_app             = "app"
	_salt            = "salt"
	_token           = "token"
	_disregard       = "disregard"
	_key             = "key"
	_server          = "server"
	_permissionCheck = "permissionCheck"
	_path            = "path"
	_method          = "method"
)

const (
	_data     = "data"
	_metadata = "metaData"
	_user     = "user"
	_role     = "role"
)

type Auth struct {
	appId  string
	config *viper.Viper
}

func New(appId string, config *viper.Viper) middleware.Object {
	return &Auth{
		appId:  appId,
		config: config,
	}
}

func (c *Auth) GetConfig() *viper.Viper {
	return c.config
}

func (a *Auth) SetConfig(config *viper.Viper) {
	a.config = config
}

func (a *Auth) Handle() gin.HandlerFunc {
	var (
		enable = a.config.GetBool(configer.GetConfigPath(_auth, _enable))
		//直接通過的 key（ app 專用 )
		disregardKey = a.config.GetString(configer.GetConfigPath(_auth, _disregard, _key))
		//取得 auth server 資料
		permissionApi = a.config.GetString(configer.GetConfigPath(_auth, _server)) + a.config.GetString(configer.GetConfigPath(_auth, _permissionCheck, _path))
		method        = a.config.GetString(configer.GetConfigPath(_auth, _permissionCheck, _method))
	)
	if enable {
		return func(c *gin.Context) {

			//判斷是否有直接通過的 key 在 header

			if key := c.GetHeader(strings.Join([]string{_auth, _disregard, _key}, _symbol)); key != "" {
				if key == disregardKey {
					c.Next()
					return
				}
			}

			client := &http.Client{}

			//對 permission 的 api 進行存取
			req, err := http.NewRequest(method, permissionApi, nil)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, e.UnknownError(err.Error()))
				return
			}

			//設定權限控制相關 header
			req.Header.Set(strings.Join([]string{_auth, _app}, _symbol), a.appId)
			req.Header.Set(strings.Join([]string{_auth, _token}, _symbol), c.GetHeader(strings.Join([]string{a.appId, _auth, _token}, _symbol)))
			req.Header.Set(strings.Join([]string{_auth, _salt}, _symbol), c.GetHeader(strings.Join([]string{a.appId, _auth, _salt}, _symbol)))
			req.Header.Set(strings.Join([]string{_auth, _path}, _symbol), c.Request.URL.Path)
			req.Header.Set(strings.Join([]string{_auth, _method}, _symbol), c.Request.Method)

			//送出 request
			resp, err := client.Do(req)

			if resp == nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"errMsg": "Connect to authority service failed.",
				})
				return
			}

			defer resp.Body.Close()

			//讀取 response body
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
				return
			}
			//將 response body 轉換成 map
			var response = make(map[string]interface{})
			if err := json.Unmarshal(bodyBytes, &response); err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
				return
			}
			if resp.StatusCode == http.StatusOK {
				data := response[_data].(map[string]interface{})[_metadata].([]interface{})
				for _, d := range data {
					meta := new(model.UserMeta)
					tmp, _ := json.Marshal(d)
					json.Unmarshal(tmp, &meta)
					if meta.Key == _role {
						c.Set(_user, meta)
						break
					}
				}
				c.Next()
				return
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
	} else {
		return func(c *gin.Context) {
			c.Next()
		}
	}
}
