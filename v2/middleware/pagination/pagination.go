package pagination

import (
	"github.com/codingXiang/cxgateway/v2/middleware"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
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
		data := make(map[string]interface{})
		if in, isExist := c.GetQuery(PAGE_SIZE); isExist {
			data[PAGE_SIZE] = in
		}
		if in, isExist := c.GetQuery(PAGE); isExist {
			data[PAGE] = in
		}
		c.Set(DATA, data)
		c.Next()
	}
}
