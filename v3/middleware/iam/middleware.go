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
	"net/http"
)

type Handler struct {
	config *viper.Viper
	cache  *redis.RedisClient
	url    string
}

const (
	svc         = "iam"
	skipAuthKey = "skipAuthKey"
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

	h := &Handler{
		cache: cache,
		url:   url,
	}

	h.SetConfig(config)
	return h
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

		if key := c.Request.URL.Query().Get("schedule_key"); key == h.GetConfig().GetString(configer.GetConfigPath(svc, skipAuthKey)) {
			c.Next()
			return
		}
		obj := &Object{
			Object: c.Request.RequestURI,
			Action: c.Request.Method,
		}

		if _type := c.GetHeader(PolicyType); _type != "" {
			obj.Type = _type
		} else {
			obj.Type = c.GetString(PolicyType)
		}
		resp, err := h.verify(jwt, salt, obj)
		defer fasthttp.ReleaseResponse(resp)
		if err != nil {
			response.SetResponse(c, "Can not access", nil, []string{err.Error()}, nil)
			c.AbortWithStatusJSON(response.StatusBadGateway(c))
			return
		} else {
			msg := "Can not access"
			switch resp.StatusCode() {
			case http.StatusUnauthorized:
				response.SetResponse(c, msg, nil, []string{"Auth failed, please check jwt token"}, nil)
				c.AbortWithStatusJSON(response.StatusUnauthorized(c))
				break
			case http.StatusForbidden:
				response.SetResponse(c, msg, nil, []string{"Auth failed, please check jwt token"}, nil)
				c.AbortWithStatusJSON(response.StatusForbidden(c))
				break
			default:
				result, e := Resp2User(resp.Body())
				if e != nil {
					c.Abort()
					return
				}
				c.Set("userInfo", result.Data)
				c.Next()
			}
		}

	}
}

func (h *Handler) verify(jwt, salt string, object *Object) (*fasthttp.Response, error) {
	logger.Log.Info("start get user auth")
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req) // <- do not forget to release
	//defer fasthttp.ReleaseResponse(resp) // <- do not forget to release

	req.SetRequestURI(h.url)

	req.Header.Add("Authorization", jwt)
	req.Header.Add("AuthSalt", salt)
	req.Header.Add("Accept", "application/json")
	req.Header.Add(PolicyType, object.Type)
	req.Header.Add(Endpoint, object.Object)
	req.Header.Add(Method, object.Action)

	err := fasthttp.Do(req, resp)
	if err != nil {
		return nil, err
	}

	return resp, err
}
