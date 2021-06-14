package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/codingXiang/configer/v2"
	"github.com/codingXiang/cxgateway/v3/middleware"
	"github.com/codingXiang/cxgateway/v3/util/response"
	"github.com/codingXiang/go-logger/v2"
	"github.com/codingXiang/go-orm/v2/redis"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
)

type Handler struct {
	config *viper.Viper
	cache  *redis.RedisClient
	url    string
}

const (
	userService = "userservice"
	_url        = "url"
	appId       = "appId"
	namespace   = "namespace"
)

func New(config *viper.Viper, cache *redis.RedisClient) middleware.Object {
	var (
		url        string
		authUrl    = config.GetString(configer.GetConfigPath(userService, _url))
		_appId     = config.GetString(configer.GetConfigPath(userService, appId))
		_namespace = config.GetString(configer.GetConfigPath(userService, namespace))
	)

	if _namespace != "" {
		url = fmt.Sprintf(authUrl, _appId+"."+_namespace)
	} else {
		url = fmt.Sprintf(authUrl, _appId)
	}

	return &Handler{
		cache: cache,
		url:   url,
	}
}

func (h *Handler) SetConfig(config *viper.Viper) {
	h.config = config
}

func (h *Handler) GetConfig() *viper.Viper {
	return h.config
}

func (h *Handler) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwt := c.GetHeader("Authorization")
		salt := c.GetHeader("AuthSalt")
		if jwt == "" {
			response.SetResponse(c, "header must have `Authorization`", nil, []string{"header must have `Authorization`"}, nil)
			c.Set("AuthStatus", false)
			c.AbortWithStatusJSON(response.StatusBadRequest(c))
			return
		}
		if salt == "" {
			response.SetResponse(c, "header must have `AuthSalt`", nil, []string{"header must have `AuthSalt`"}, nil)
			c.Set("AuthStatus", false)
			c.AbortWithStatusJSON(response.StatusBadRequest(c))
			return
		}
		var (
			info *User
			err  error
		)
		info, err = h.verify(jwt, salt)
		if err != nil {
			response.SetResponse(c, "JWT Auth Failed", nil, []string{err.Error()}, nil)
			c.Set("AuthStatus", false)
			c.AbortWithStatusJSON(response.StatusUnauthorized(c))
			return
		}
		c.Set("AuthStatus", true)
		c.Set(UserInfo, info)
		//c.Set(response.ModuleName, h.GetConfig().Get(response.ModuleName))
		c.Next()
	}
}

func (h *Handler) verify(jwt, salt string) (*User, error) {
	logger.Log.Info("start get user auth")
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)   // <- do not forget to release
	defer fasthttp.ReleaseResponse(resp) // <- do not forget to release

	req.SetRequestURI(h.url)

	req.Header.Add("Authorization", jwt)
	req.Header.Add("AuthSalt", salt)
	req.Header.Add("Accept", "application/json")
	//resp, err := http.DefaultClient.Do(req)

	err := fasthttp.Do(req, resp)

	logger.Log.Debug("end get user auth")
	if err != nil {
		return nil, err
	}
	logger.Log.Debug("[Auth] url = ", h.url, ", token = ", jwt, ", salt = ", salt)
	//defer resp.Body.Close()

	if resp.StatusCode() > 399 {
		err = errors.New("Auth failed, please check jwt token or salt")
		return nil, err
	}

	return getInfo(resp.Body())
}

func getInfo(in []byte) (*User, error) {
	_resp := new(response.Response)
	info := new(Response)

	if err := json.Unmarshal(in, &_resp); err != nil {
		return nil, err
	}

	out, err := json.Marshal(_resp.Data)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(out, &info); err != nil {
		return nil, err
	}
	return info.User, nil
}
