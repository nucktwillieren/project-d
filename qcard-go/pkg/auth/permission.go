package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nucktwillieren/project-d/qcard-go/pkg/utils"
)

func JwtRolePermissionMiddleware(roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if role, ok := c.Get("auth-role"); ok {
			switch v := role.(type) {
			case string:
				if !utils.StrIn(v, roles) {
					c.JSON(http.StatusUnauthorized, gin.H{
						"error": "Role Forbbiden",
					})
					c.Abort()
				}
			}
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Role Forbbiden",
			})
			c.Abort()
		}
	}
}

func JwtAnonymousUserForbbiden(anonymousUser string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if username, ok := GetUser(c); !ok || username == anonymousUser {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Anonymous Forbbiden",
			})
			c.Abort()
		}
	}
}
