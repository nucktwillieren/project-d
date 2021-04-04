package handler

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nucktwillieren/project-d/qcard-go/internal/models"
	"github.com/nucktwillieren/project-d/qcard-go/pkg/auth"
	"github.com/nucktwillieren/project-d/qcard-go/pkg/db"
)

func GetPair(c *gin.Context) {
	var userModel models.User
	db := c.MustGet("db-params").(db.DBParams)

	db.PG.Model(&userModel).Where("username = ?", c.Param("username")).Relation("Friends").Relation("Pairing").Select()
	if userModel.Username == auth.GetUsername(c) {
		log.Println(userModel)
	}
}
