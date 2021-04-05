package xlimit

import (
	"context"
	"log"
	"net/http"
	"time"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nucktwillieren/project-d/qcard-go/pkg/auth"
	"github.com/nucktwillieren/project-d/qcard-go/pkg/utils"
	"google.golang.org/grpc"
)

func XlimitMiddleware(conn *grpc.ClientConn, identity string, increase uint64) gin.HandlerFunc {
	return func(c *gin.Context) {
		client := NewXLimitClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		r, err := client.CheckAndIncrease(ctx, &XLimitCheckRequest{Identity: identity, IncreaseNumber: increase})
		if err != nil && err != LimitExceedError {
			log.Fatalf("could not check: %v", err)
		}
		if err != LimitExceedError {
			log.Printf("Check %s: %v(Remain:%v)(Reset:%v)", r.GetIdentity(), r.GetIsAllowed(), r.GetCountRemaining(), r.GetTimeleft())
			c.Header("X-RateLimit-Remaining", strconv.Itoa(int(r.GetCountRemaining())))
			c.Header("X-RateLimit-Reset", strconv.Itoa(int(r.GetTimeleft())))
		} else {
			c.JSON(http.StatusTooManyRequests, gin.H{"err": "Too Many Requests"})
			c.Abort()
		}
	}
}

func XlimitMiddlewareWithIPAndUser(conn *grpc.ClientConn) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := utils.GetRealIPFromGateway(c)
		realIP := ""

		if ip == "" {
			realIP = c.ClientIP()
		} else {
			realIP = ip
		}
		identityWithIP := realIP + ":" + c.Request.URL.Path + ":" + c.Request.Method + ":" + auth.GetUsername(c)

		client := NewXLimitClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

		defer cancel()
		r, err := client.CheckAndIncrease(ctx, &XLimitCheckRequest{Identity: identityWithIP, IncreaseNumber: 1})
		if err != nil && err != LimitExceedError {
			log.Fatalf("could not check: %v", err)
		}
		if err != LimitExceedError {
			log.Printf("Check %s: %v(Remain:%v)(Reset:%v)", r.GetIdentity(), r.GetIsAllowed(), r.GetCountRemaining(), r.GetTimeleft())
			c.Header("X-RateLimit-Remaining", strconv.Itoa(int(r.GetCountRemaining())))
			c.Header("X-RateLimit-Reset", strconv.Itoa(int(r.GetTimeleft())))
		} else {
			c.JSON(http.StatusTooManyRequests, gin.H{"err": "Too Many Requests"})
			c.Abort()
		}
	}
}
