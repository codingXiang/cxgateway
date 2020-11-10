package delivery

import "github.com/gin-gonic/gin"

type HttpHandler interface {
	GetServerVersion(c *gin.Context) error
	CheckVersion(c *gin.Context) error
	Upgrade(c *gin.Context) error
}
