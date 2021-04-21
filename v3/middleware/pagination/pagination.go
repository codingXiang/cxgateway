package pagination

import (
	"errors"
	"github.com/codingXiang/cxgateway/v3/constants"
	"github.com/codingXiang/cxgateway/v3/middleware"
	"github.com/codingXiang/cxgateway/v3/util/response"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"strconv"
)



var defaultDatas = []*defaultData{newDefaultData(constants.PageSize, 10), newDefaultData(constants.Page, 1)}

type Handler struct{}

type defaultData struct {
	Key   string
	Value int
}

func newDefaultData(key string, value int) *defaultData {
	return &defaultData{
		key,
		value,
	}
}

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
		statusCode, exception := response.StatusBadRequest(c)
		for _, data := range defaultDatas {
			if val, err := getData(c, data); err == nil {
				c.Set(data.Key, val)
			} else {
				exception.Message = err.Error()
				c.AbortWithStatusJSON(statusCode, exception)
				return
			}
		}
		c.Next()
	}
}

func getData(c *gin.Context, param *defaultData) (int, error) {
	if in, isExist := c.GetQuery(param.Key); isExist {
		if out, err := isInt(in); err != nil {
			return 0, errors.New(param.Key + " must to be int")
		} else {
			return out, nil
		}
	} else {
		return param.Value, nil
	}
}

func isInt(in string) (int, error) {
	res, err := strconv.Atoi(in)
	return res, err
}
