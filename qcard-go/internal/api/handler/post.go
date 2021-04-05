package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nucktwillieren/project-d/qcard-go/internal/models"
	"github.com/nucktwillieren/project-d/qcard-go/pkg/auth"
	"github.com/nucktwillieren/project-d/qcard-go/pkg/db"
)

type GetPostControl struct {
	Number uint
}

func CreateCategory(c *gin.Context) {
	var categoryModel models.Category
	if err := c.ShouldBindJSON(&categoryModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err})
		return
	}

	db := c.MustGet("db-params").(db.DBParams)
	categoryModel.CreatedAt = time.Now().UTC()

	if res, err := db.PG.Model(&categoryModel).Insert(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err, "res": res})
		return
	} else {
		c.JSON(http.StatusOK, res)
		return
	}
}

func GetAllCategory(c *gin.Context) {
	var categoryModel []models.Category

	db := c.MustGet("db-params").(db.DBParams)

	if err := db.PG.Model(&categoryModel).Select(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err})
		return
	}
	c.JSON(http.StatusOK, categoryModel)
}

func CreatePost(c *gin.Context) {
	var postModel models.Post
	var userModel models.User
	var categoryModel models.Category

	username := auth.GetUsername(c)
	if username == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"err": "cannot get user info"})
		return
	}

	db := c.MustGet("db-params").(db.DBParams)
	if err := db.PG.Model(&userModel).Where("username = ?", username).Select(); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"err": "no user"})
		return
	}

	if err := c.ShouldBindJSON(&postModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err})
		return
	}

	if err := db.PG.Model(&categoryModel).Where("name = ?", c.Param("category_name")).Select(); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"err": "no category"})
		return
	}

	postModel.CreatedAt = time.Now().UTC()
	postModel.CreatorID = userModel.ID
	postModel.CategoryID = categoryModel.ID

	if res, err := db.PG.Model(&postModel).Insert(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"res": res, "err": err})
		return
	}

	c.JSON(http.StatusOK, postModel)
}

func GetPost(c *gin.Context) {
	var postModel []models.Post
	var control GetPostControl

	if err := c.ShouldBindQuery(&control); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err})
		return
	}

	db := c.MustGet("db-params").(db.DBParams)

	if err := db.PG.Model(&postModel).Relation("Creator").Relation("Category").Limit(int(control.Number)).Select(); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"err": err})
		return
	}

	c.JSON(http.StatusOK, postModel)
}
