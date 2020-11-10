package middleware

import (
	"github.com/gin-gonic/gin"
)

type Object interface {
	Handle() gin.HandlerFunc
}
