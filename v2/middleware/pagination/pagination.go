package pagination

import (
	"errors"
	"github.com/codingXiang/cxgateway/v2/middleware"
	"github.com/codingXiang/cxgateway/v2/util/e"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"strconv"
)

const (
	DATA      = "data"
	PAGE      = "page"
	PAGE_SIZE = "pageSize"
)

type Handler struct{}

func New() middleware.Object {
	return new(Handler)
}

func (*Handler) SetConfig(config *viper.Viper) {

}

func (c *Handler) GetConfig() *viper.Viper {
	return viper.New()
}

func (c *Handler) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		exceptionMsg := &e.APIException{
			Code:      http.StatusBadRequest,
			ErrorCode: http.StatusBadRequest,
			Message:   "",
		}
		data := make(map[string]interface{})
		if d, err := setData(c, data, PAGE_SIZE); err == nil {
			data = d
		} else {
			exceptionMsg.Message = err.Error()
			c.AbortWithStatusJSON(http.StatusBadRequest, exceptionMsg)
			return
		}
		if d, err := setData(c, data, PAGE); err == nil {
			data = d
		} else {
			exceptionMsg.Message = err.Error()
			c.AbortWithStatusJSON(http.StatusBadRequest, exceptionMsg)
			return
		}
		c.Set(DATA, data)
		c.Next()
	}
}

func setData(c *gin.Context, data map[string]interface{}, key string) (map[string]interface{}, error) {
	if in, isExist := c.GetQuery(key); isExist {
		if isInt(in) != nil {
			return data, errors.New(key + " must to be int")
		}
		data[key] = in
	}
	return data, nil
}

func isInt(in string) error {
	_, err := strconv.Atoi(in)
	return err
}
