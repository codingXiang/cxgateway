package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Object interface {
	GetConfig() *viper.Viper
	SetConfig(config *viper.Viper)
	Handle() gin.HandlerFunc
}
