package auth

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Claims struct {
	Account string `json:"account"`
	Role    string `json:"role"`
	Scope   string `json:"scope"`
	jwt.StandardClaims
}

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

func SetBaseParams(params JWTParams) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("jwt-params", params)
	}
}

func NewJwt(account string, role string, scope string, aud string, exp int64, issuer string, nbf int64, subject string, secret []byte) string {
	now := time.Now()
	jwtID := account + strconv.FormatInt(now.Unix(), 10)

	claims := Claims{
		Account: account,
		Role:    role,
		Scope:   scope,
		StandardClaims: jwt.StandardClaims{
			Audience:  aud,
			ExpiresAt: exp,
			Id:        jwtID,
			IssuedAt:  now.Unix(),
			Issuer:    issuer,
			NotBefore: nbf,
			Subject:   subject,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(secret)

	if err != nil {
		fmt.Println(err)
	}

	return token
}

func JwtVerified(token string, secret []byte) (bool, string, *Claims) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return secret, nil
	})
	message := "token is valid."

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				message = "token is malformed"
			} else if ve.Errors&jwt.ValidationErrorUnverifiable != 0 {
				message = "token could not be verified because of signing problems"
			} else if ve.Errors&jwt.ValidationErrorSignatureInvalid != 0 {
				message = "signature validation failed"
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				message = "token is expired"
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				message = "token is not yet valid before sometime"
			} else {
				message = "can not handle this token"
			}
		}
		return false, message, &Claims{}
	}

	if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
		return true, message, claims
	}

	return false, message, &Claims{}
}

func GetTokenFromHeader(c *gin.Context, headerKey string, tokenType string) string {
	headerValue := c.GetHeader(headerKey)
	if len(headerValue) != 0 {
		if v := strings.Split(headerValue, tokenType+" "); len(v) > 1 {
			return v[1]
		}
	}
	return ""
}

func JwtAuthMiddleware(headerKey string, tokenType string, secret []byte, anonymousUser string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if token := GetTokenFromHeader(c, headerKey, tokenType); token != "" {
			if ok, msg, claims := JwtVerified(token, secret); ok {
				if tempClaims, e := c.Get("claims"); e {
					switch v := tempClaims.(type) {
					case map[string]*Claims:
						v[claims.Scope] = claims
						c.Set("claims", v)
						c.Set("auth-user", claims.Account)
						c.Set("auth-role", claims.Role)
					default:
						c.Set("claims", map[string]*Claims{claims.Scope: claims})
						c.Set("auth-user", claims.Account)
						c.Set("auth-role", claims.Role)
					}
				} else {
					c.Set("claims", map[string]*Claims{claims.Scope: claims})
					c.Set("auth-user", claims.Account)
					c.Set("auth-role", claims.Role)
				}
				return
			} else {
				c.Set("auth-user", anonymousUser)
				c.Set("auth-user-unknown-msg", msg)
				return
			}
		}
		c.Set("auth-user", anonymousUser)
	}
}

func GetUser(c *gin.Context) (string, bool) {
	if user, ok := c.Get("auth-user"); ok {
		switch v := user.(type) {
		case string:
			return v, true
		}
	}
	return "", false
}

func GetUsername(c *gin.Context) string {
	if user, ok := c.Get("auth-user"); ok {
		switch v := user.(type) {
		case string:
			return v
		}
	}
	return ""
}

func GetClaims(c *gin.Context) (map[string]*Claims, bool) {
	if claims, ok := c.Get("claims"); ok {
		switch v := claims.(type) {
		case map[string]*Claims:
			return v, true
		}
	}
	return map[string]*Claims{}, false
}
