package handler

import (
	"errors"
	"log"
	"net"
	"net/http"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/nucktwillieren/project-d/qcard-go/pkg/auth"
)

var (
	testU            = "test"
	testP            = "test"
	TypePointerError = errors.New("Arguement should be pointer")
	emailRegex       = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

func IsEmailValid(e string) bool {
	if len(e) < 3 && len(e) > 254 {
		return false
	}
	if !emailRegex.MatchString(e) {
		return false
	}
	parts := strings.Split(e, "@")
	mx, err := net.LookupMX(parts[1])
	if err != nil || len(mx) == 0 {
		return false
	}
	return true
}

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegistrationPayload struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	PasswordSecond string `json:"password_second"`
	Email          string `json:"email"`
}

func GetLoginToken(c *gin.Context, username string) {
	parms := c.MustGet("jwt-params").(auth.JWTParams)
	token := auth.NewJwt(
		username,
		"", "", "",
		time.Now().Add(time.Second*3600).Unix(),
		"",
		time.Now().Unix(),
		"",
		parms.Secret,
	)

	c.JSON(http.StatusOK, gin.H{
		"user":         username,
		"token_type":   "Bearer",
		"access_token": token,
		"expires_in":   3600,
	})
}

func Login(c *gin.Context) {
	var credential Credential
	if err := c.ShouldBindJSON(&credential); err != nil {
		c.JSON(http.StatusBadRequest, "Wrong Format")
		return
	}

	if credential.Password == testP && credential.Username == testU {
		GetLoginToken(c, credential.Username)
	}
}

func PointerCheck(v interface{}) bool {
	return reflect.ValueOf(v).Kind() == reflect.Ptr
}

func CheckUsernameDuplicated(db *pg.DB, model interface{}, username string) (bool, error) {
	if PointerCheck(model) {
		return db.Model(model).Where("username = ?", username).Exists()
	}
	log.Panicln(TypePointerError)
	return true, TypePointerError
}

func Registration(c *gin.Context) {
	var payload RegistrationPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, "Wrong Format")
		return
	}
	//if exists, err := CheckUsernameDuplicated(){
	//
	//}
	GetLoginToken(c, payload.Username)
}
