package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"strings"
)

//RequestIDMiddleware 中間件回應Header
func RequestIDMiddleware(appID string) gin.HandlerFunc {
	return func(c *gin.Context) {
		trackID := uuid.NewV4()
		trackKey := fmt.Sprintf("X-%s-Track-Id", strings.ToUpper(appID))
		c.Writer.Header().Set(trackKey, trackID.String())
		c.Next()
	}
}

//RequestVersion
func RequestVersion(version string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("app-version", version)
		c.Next()
	}
}