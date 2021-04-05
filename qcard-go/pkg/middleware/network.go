package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/nucktwillieren/project-d/qcard-go/pkg/utils"
)

func SetRealIPFromGateway(c *gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("real-ip", utils.GetRealIPFromGateway(c))
	}
}
