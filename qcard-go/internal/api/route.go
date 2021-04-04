package api

import (
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"github.com/nucktwillieren/project-d/qcard-go/internal/api/handler"
	"github.com/nucktwillieren/project-d/qcard-go/pkg/auth"
	"github.com/nucktwillieren/project-d/qcard-go/pkg/db"
)

var (
	dbMap map[string]*pg.DB
)

func init() {
	dbMap = db.YamlToPGOptions(os.Getenv("QCARD_GO_DB_CONFIG_PATH"))
}

func Setup() *gin.Engine {
	RouterBase := gin.Default()
	secret := []byte(uuid.New().String())
	RouterBase.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"}, // In dev, allow all.
		AllowMethods: []string{"GET", "POST", "HEAD", "PUT", "PATCH"},
		AllowHeaders: []string{"Authorization", "Content-Type", "Upgrade", "Origin",
			"Connection", "Accept-Encoding", "Accept-Language", "Host", "Access-Control-Request-Method", "Access-Control-Request-Headers"},
		ExposeHeaders: []string{"Authorization", "Content-Type", "Upgrade", "Origin",
			"Connection", "Accept-Encoding", "Accept-Language", "Host", "Access-Control-Request-Method", "Access-Control-Request-Headers"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	v1 := RouterBase.Group("api/v1/")
	{
		authGroup := v1.Group("auth/")
		authGroup.Use(auth.SetJWTParams(auth.JWTParams{Secret: secret}))
		authGroup.Use(db.SetDBParams(db.DBParams{PG: dbMap["qcard"]}))
		{
			authGroup.POST("login", handler.Login)
			authGroup.POST("registration", handler.Registration)
		}

		user := v1.Group("user/")
		user.Use(db.SetDBParams(db.DBParams{PG: dbMap["qcard"]}))
		{

		}

		//admin := v1.Group("admin/")
	}

	return RouterBase
}
