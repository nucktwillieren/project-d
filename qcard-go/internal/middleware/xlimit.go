package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/nucktwillieren/project-d/qcard-go/internal/xlimit"
)

func XlimitClient(client *xlimit.XLimitClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("x-limit-client", client)
	}
}

func XlimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
