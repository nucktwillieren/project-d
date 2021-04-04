package db

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
)

type DBParams struct {
	Redis *redis.Client
	PG    *pg.DB
	Mongo *mongo.Client
}

func SetDBParams(params DBParams) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db-params", params)
	}
}
