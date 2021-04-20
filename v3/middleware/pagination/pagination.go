package pagination

import (
	"errors"
	"github.com/codingXiang/cxgateway/v3/middleware"
	"github.com/codingXiang/cxgateway/v3/util/response"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"strconv"
)

const (
	PageParameter = "pageParameter" //傳遞參數 key
	Page          = "page"          //分頁
	PageSize      = "limit"         //每頁資料限制筆數
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
		statusCode, exception := response.StatusBadRequest(c)
		data := make(map[string]interface{})
		for _, key := range []string{PageSize, Page} {
			if d, err := setData(c, data, key); err == nil {
				data = d
			} else {
				exception.Message = err.Error()
				c.AbortWithStatusJSON(statusCode, exception)
				return
			}
		}
		c.Set(PageParameter, data)
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
