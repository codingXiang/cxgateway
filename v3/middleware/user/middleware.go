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
	"time"
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
		if jwt == "" {
			response.SetResponse(c, "header must have `Authorization`", nil, []string{"header must have `Authorization`"}, nil)
			c.Set("AuthStatus", false)
			c.AbortWithStatusJSON(response.StatusBadRequest(c))
			return
		}
		var (
			info *User
			err  error
		)
		if info, err = h.getCache(jwt); err == nil {
			logger.Log.Debug("Get jwt cache from redis success, key = ", jwt)
		} else {
			info, err = h.verify(jwt)
			if err == nil {
				if err = h.setCache(jwt, info); err == nil {
					logger.Log.Debug("Set jwt cache to redis success, key = ", jwt)
				} else {
					response.SetResponse(c, "Set cache failed", nil, []string{err.Error()}, nil)
					c.AbortWithStatusJSON(response.StatusUnauthorized(c))
					c.Set("AuthStatus", false)
					return
				}
			} else {
				response.SetResponse(c, "JWT Auth Failed", nil, []string{err.Error()}, nil)
				c.Set("AuthStatus", false)
				c.AbortWithStatusJSON(response.StatusUnauthorized(c))
				return
			}
		}

		c.Set(UserInfo, info)
		c.Set(response.ModuleName, h.GetConfig().Get(response.ModuleName))
		c.Next()
	}
}

func (h *Handler) verify(jwt string) (*User, error) {
	logger.Log.Info("start get user auth")
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)   // <- do not forget to release
	defer fasthttp.ReleaseResponse(resp) // <- do not forget to release

	req.SetRequestURI(h.url)

	//req, err := http.NewRequest(http.MethodGet, h.url, nil)
	//if err != nil {
	//	return nil, err
	//}
	req.Header.Add("Authorization", jwt)
	req.Header.Add("Accept", "application/json")
	//resp, err := http.DefaultClient.Do(req)

	err := fasthttp.Do(req, resp)

	logger.Log.Info("end get user auth")
	if err != nil {
		return nil, err
	}
	logger.Log.Debug("[Auth] url = ", h.url, ", token = ", jwt)
	//defer resp.Body.Close()

	if resp.StatusCode() > 399 {
		err = errors.New("Auth failed, please check jwt token")
		return nil, err
	}

	//if resp.StatusCode > 399 {
	//	err = errors.New("Auth failed, please check jwt token")
	//	return nil, err
	//}

	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	return nil, err
	//}
	return getInfo(resp.Body())
}

func (h *Handler) setCache(jwt string, user *User) error {
	b, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return h.cache.SetKeyValue(jwt, b, 1*time.Minute)
}

func (h *Handler) getCache(jwt string) (*User, error) {
	user := new(User)
	info, err := h.cache.GetValue(jwt)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(info), &user)
	if err != nil {
		logger.Log.Error(err.Error())
		return nil, err
	}
	return user, nil
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
