package xlimit

import (
	"context"
	"log"
	"net/http"
	"time"

	"strconv"

	"github.com/gin-gonic/gin"
)

func XlimitMiddleware(client *xLimitClient, identity string, increase uint64) gin.HandlerFunc {
	return func(c *gin.Context) {
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
