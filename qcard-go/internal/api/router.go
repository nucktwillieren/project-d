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
	"github.com/nucktwillieren/project-d/qcard-go/pkg/xlimit"
	"google.golang.org/grpc"
)

var (
	dbMap      map[string]*pg.DB
	clientConn *grpc.ClientConn
)

func init() {
	dbMap = db.YamlToPGOptions(os.Getenv("QCARD_GO_DB_CONFIG_PATH"))
	clientConn = xlimit.NewClientConn((os.Getenv("XLIMIT_GRPC_ADDR")))
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
	v1.Use(db.SetDBParams(db.DBParams{PG: dbMap["qcard"]}))
	{
		authGroup := v1.Group("auth/")
		authGroup.Use(auth.SetJWTParams(auth.JWTParams{Secret: secret}))
		{
			authGroup.POST("login", handler.Login)
			authGroup.POST("registration", handler.Registration)
		}

		jwtProtectedGroup := v1.Group("")
		jwtProtectedGroup.Use(auth.JwtAuthMiddleware("Authorization", "Bearer", secret, "unknown"))
		jwtProtectedGroup.Use(auth.JwtAnonymousUserForbbiden("unknown"))
		{
			user := jwtProtectedGroup.Group("user/")
			{
				user.GET(":username", handler.GetUser)
			}

			post := jwtProtectedGroup.Group("post/")
			{
				post.POST("", handler.CreatePost)
				post.GET(":id")
			}

			category := jwtProtectedGroup.Group("category/")
			{
				category.POST("", handler.CreateCategory)
			}

			pair := jwtProtectedGroup.Group("pair/")
			pair.Use(xlimit.XlimitMiddlewareWithIPAndUser(clientConn))
			{
				pair.GET("", handler.GetAllPair)
				pair.GET(":username", handler.GetPair)
				pair.POST(":username", handler.CreateRandomPair)
				pair.PATCH("null", handler.SetPairingNull)
				pair.DELETE("", handler.CleanPair)
			}
		}

		noAuthGroup := v1.Group("")
		{
			post := noAuthGroup.Group("category/")
			{
				post.GET("", handler.GetAllCategory)
			}
		}

		//admin := v1.Group("admin/")
	}

	return RouterBase
}
