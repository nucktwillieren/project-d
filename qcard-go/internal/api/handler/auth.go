package handler

import (
	"errors"
	"log"
	"net"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/nucktwillieren/project-d/qcard-go/internal/models"
	"github.com/nucktwillieren/project-d/qcard-go/pkg/auth"
	"github.com/nucktwillieren/project-d/qcard-go/pkg/db"
	"github.com/nucktwillieren/project-d/qcard-go/pkg/utils"
)

var (
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
		c.JSON(http.StatusBadRequest, gin.H{"err": "Wrong Format"})
		return
	}

	db := c.MustGet("db-params").(db.DBParams)
	var userModel models.User

	if err := db.PG.Model(&userModel).Where("username = ?", credential.Username).Select(); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"err": "No User"})
		return
	}

	if ok, err := auth.SaltingVerify(credential.Password, userModel.Password); ok && err == nil {
		GetLoginToken(c, userModel.Username)
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"err": "Wrong Format"})
	//GetLoginToken(c, credential.Username)
}

func CheckDuplicated(db *pg.DB, model interface{}, key string, value string) (bool, error) {
	if utils.PointerCheck(model) {
		return db.Model(model).Where(key+" = ?", value).Exists()
	}
	log.Panicln(TypePointerError)
	return true, TypePointerError
}

func CheckSecondPassword(pw1 string, pw2 string) bool {
	return pw1 == pw2
}

func Registration(c *gin.Context) {
	db := c.MustGet("db-params").(db.DBParams)
	var payload RegistrationPayload
	var userModel models.User

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Wrong Format"})
		return
	}

	if !CheckSecondPassword(payload.Password, payload.PasswordSecond) {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Second Password Checking Failed"})
		return
	}

	if exists, _ := CheckDuplicated(db.PG, &userModel, "username", payload.Username); exists {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Username Duplicated"})
		return
	}

	if exists, _ := CheckDuplicated(db.PG, &userModel, "email", payload.Email); exists {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Email Duplicated"})
		return
	}

	if saltedPassword, err := auth.Salting(payload.Password, 0, ""); err == nil {
		// iteration = 0 => will get default = 180000
		// salt = "" => will get default, will call the random_string funciton
		userModel.Username = payload.Username
		userModel.Email = payload.Email
		userModel.Password = saltedPassword
		if res, err := db.PG.Model(&userModel).Insert(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": "Create User Error"})
			log.Println(res, err)
			return
		}

		GetLoginToken(c, payload.Username)
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"err": "Unknown Error"})
}
