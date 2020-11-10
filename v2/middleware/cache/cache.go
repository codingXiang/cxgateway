package cache

import (
	"github.com/codingXiang/cxgateway/v2/middleware"
	"github.com/gin-gonic/gin"
	"strconv"
)

const (
	Parameter = "realtime"
	Key       = "enableCache"
)

type Cache struct{}

func New() middleware.Object {
	return new(Cache)
}

func (c *Cache) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			enableCache = true
		)
		if tmp, isExist := c.GetQuery(Parameter); isExist {
			if cache, err := strconv.ParseBool(tmp); err == nil {
				enableCache = !cache
			}
		}
		c.Set(Key, enableCache)
		c.Next()
	}
}
