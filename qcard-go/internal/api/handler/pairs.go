package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nucktwillieren/project-d/qcard-go/internal/models"
	"github.com/nucktwillieren/project-d/qcard-go/pkg/auth"
	"github.com/nucktwillieren/project-d/qcard-go/pkg/db"
)

func GetPair(c *gin.Context) {
	var userModel models.User
	var pairUser models.User
	db := c.MustGet("db-params").(db.DBParams)

	db.PG.Model(&userModel).Where("username = ?", c.Param("username")).Relation("Pairing").Select()
	if userModel.Username != auth.GetUsername(c) {
		c.JSON(http.StatusNotFound, gin.H{"err": "no user"})
		return
	}
	if err := db.PG.Model(&pairUser).
		Where("id = ?", userModel.Pairing.UserTwoID).
		Select(); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"err": "no user"})
		return
	}

	c.JSON(http.StatusOK, pairUser)
}

func GetAllPair(c *gin.Context) {
	var pairs []models.Friend
	db := c.MustGet("db-params").(db.DBParams)
	db.PG.Model(&pairs).Select()
	c.JSON(http.StatusOK, pairs)
}

func CreateRandomPair(c *gin.Context) {
	var friendModelOne models.Friend
	var friendModelTwo models.Friend
	var currentUser models.User
	var randomUser models.User

	db := c.MustGet("db-params").(db.DBParams)

	db.PG.Model(&currentUser).Where("username = ?", c.Param("username")).Relation("Friends").Relation("Pairing").Select()
	if currentUser.Username != auth.GetUsername(c) {
		c.JSON(http.StatusNotFound, gin.H{"err": "no user"})
		return
	}

	if currentUser.Pairing != nil {
		c.JSON(http.StatusForbidden, gin.H{"err": "already having pair"})
		return
	}

	if err := db.PG.Model(&randomUser).
		Where("id != ?", currentUser.ID).
		Where(`not exists (SELECT 1 FROM friends WHERE user_one_id = ? and user_two_id = "user"."id")`, currentUser.ID).
		OrderExpr("random()").Limit(1).Select(); err != nil {

		c.JSON(http.StatusNotFound, gin.H{"err": "no user"})
		return
	}

	friendModelOne.UserOneID = currentUser.ID
	friendModelOne.UserTwoID = randomUser.ID
	friendModelOne.Pair = false
	friendModelOne.CreatedAt = time.Now().UTC()

	friendModelTwo.UserOneID = randomUser.ID
	friendModelTwo.Pair = false
	friendModelTwo.UserTwoID = currentUser.ID
	friendModelTwo.CreatedAt = time.Now().UTC()

	db.PG.Model(&friendModelOne).Insert()
	db.PG.Model(&friendModelTwo).Insert()

	currentUser.PairingID = friendModelOne.ID
	randomUser.PairingID = friendModelTwo.ID
	currentUser.Pairing = &friendModelOne
	randomUser.Pairing = &friendModelTwo

	log.Println(currentUser.PairingID, randomUser.PairingID)
	if res, err := db.PG.Model(&currentUser).WherePK().Column("pairing_id").Update(); err != nil {
		log.Println(res, err)
		return
	}
	if res, err := db.PG.Model(&randomUser).WherePK().Column("pairing_id").Update(); err != nil {
		log.Println(res, err)
		return
	}

	c.JSON(http.StatusOK, randomUser)
}

func ConfirmPair(c *gin.Context) {

}

func SetPairingNull(c *gin.Context) {
	rawQuery := `UPDATE public.users SET pairing_id=null;`
	db := c.MustGet("db-params").(db.DBParams)
	res, err := db.PG.Exec(rawQuery)
	c.JSON(http.StatusOK, gin.H{"res": res, "err": err})
}

func CleanPair(c *gin.Context) {
	rawQuery := `DELETE FROM public.friends;`
	db := c.MustGet("db-params").(db.DBParams)
	res, err := db.PG.Exec(rawQuery)
	c.JSON(http.StatusOK, gin.H{"res": res, "err": err})
}
