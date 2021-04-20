package service

import (
	"errors"
	"github.com/codingXiang/cxgateway/v3/middleware"
	"github.com/codingXiang/cxgateway/v3/util/response"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"strconv"
)

type Handler struct {
	config *viper.Viper
}

func New(moduleCode response.ModuleCode) middleware.Object {
	config := viper.New()
	config.Set(response.ModuleName, moduleCode)
	return &Handler{
		config: config,
	}
}

func (*Handler) SetConfig(config *viper.Viper) {

}

func (c *Handler) GetConfig() *viper.Viper {
	return c.config
}

func (h *Handler) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(response.ModuleName, h.GetConfig().Get(response.ModuleName))
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
