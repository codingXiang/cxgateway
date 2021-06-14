package iam

import (
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
	svc         = "iam"
	_url        = "url"
	_endpoint   = "endpoint"
	_permission = "permission"
	appId       = "appId"
	namespace   = "namespace"
)

const (
	PolicyType = "IAM-Policy-Type"
	Endpoint   = "IAM-Access-Endpoint"
	Method     = "IAM-Access-Method"
)

func New(config *viper.Viper, cache *redis.RedisClient) middleware.Object {
	var (
		url        string
		authUrl    = config.GetString(configer.GetConfigPath(svc, _url))
		_appId     = config.GetString(configer.GetConfigPath(svc, appId))
		_namespace = config.GetString(configer.GetConfigPath(svc, namespace))
		e          = config.GetString(configer.GetConfigPath(svc, _endpoint, _permission))
	)

	url = authUrl + e

	if _namespace != "" {
		url = fmt.Sprintf(url, _appId+"."+_namespace)
	} else {
		url = fmt.Sprintf(url, _appId)
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
		obj := &Object{
			Type:   c.GetHeader(PolicyType),
			Object: c.Request.RequestURI,
			Action: c.Request.Method,
		}

		if statusCode, err := h.verify(jwt, salt, obj); err != nil {
			response.SetResponse(c, "Can not access", nil, []string{err.Error()}, nil)
			if statusCode == 401 {
				c.AbortWithStatusJSON(response.StatusUnauthorized(c))
				return
			}
			if statusCode == 403 {
				c.AbortWithStatusJSON(response.StatusForbidden(c))
				return
			}
			if statusCode > 399 {
				c.AbortWithStatusJSON(response.StatusInternalServerError(c))
				return
			}
			return
		}
		c.Next()
	}
}

func (h *Handler) verify(jwt, salt string, object *Object) (int, error) {
	logger.Log.Info("start get user auth")
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)   // <- do not forget to release
	defer fasthttp.ReleaseResponse(resp) // <- do not forget to release

	req.SetRequestURI(h.url)

	req.Header.Add("Authorization", jwt)
	req.Header.Add("AuthSalt", salt)
	req.Header.Add("Accept", "application/json")
	req.Header.Add(PolicyType, object.Type)
	req.Header.Add(Endpoint, object.Object)
	req.Header.Add(Method, object.Action)

	err := fasthttp.Do(req, resp)
	if err != nil {
		return 500, err
	}

	if resp.StatusCode() == 401 {
		err = fmt.Errorf("Auth failed, please check jwt token")
	} else if resp.StatusCode() == 403 {
		err = fmt.Errorf("have no permission")
	} else if resp.StatusCode() > 399 {
		err = fmt.Errorf("unknown error")
	}
	return resp.StatusCode(), err
}
