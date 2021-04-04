package auth

import "github.com/gin-gonic/gin"

type JWTParams struct {
	Account string
	Role    string
	Scope   string
	Aud     string
	Exp     int64
	Issur   string
	Nbf     int64
	Subject string
	Secret  []byte
}

func SetJWTParams(params JWTParams) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("jwt-params", params)
	}
}
