package request

import "github.com/gin-gonic/gin"

func GetValues(c *gin.Context, data map[string]interface{}, params ...string) map[string]interface{} {
	if data == nil {
		data = make(map[string]interface{})
	}
	for _, p := range params {
		if in, isExist := c.Get(p); isExist {
			data[p] = in
		}
	}
	return data
}
