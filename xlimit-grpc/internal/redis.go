package internal

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/kelseyhightower/envconfig"
	"github.com/nucktwillieren/project-d/xlimit-grpc/pkg/xlimit"
)

type XlimitRedisLayerOptions struct {
	Prefix   string
	LimitNum uint64
}

type XlimitRedisLayer struct {
	client *redis.Client
	*XlimitRedisLayerOptions
}

/*
structure in redis
key => service_name:identity.
value_field: count(uint).
init_expiration: 1 hour.
*/
func (xrl *XlimitRedisLayer) CheckAndIncrease(ctx context.Context, in *xlimit.XLimitCheckRequest, out *xlimit.XLimitCheckReply) (string, *xlimit.XLimitCheckReply, error) {
	key := xrl.Prefix + ":" + in.GetIdentity()
	val, err := xrl.client.Get(ctx, key).Uint64()
	increase := in.GetIncreaseNumber()
	out.Identity = in.Identity
	switch {
	case err == redis.Nil || val == 0:
		res, setErr := xrl.client.Set(ctx, key, increase, time.Hour*1).Result()
		out.IsAllowed = true
		out.Timeleft = uint64(time.Hour) * 1
		out.CountRemaining = xrl.LimitNum - increase
		log.Printf("Create New Identify(GRPC):%v -> Key(Redis):%v, Result:%v(Err:%v)", in.GetIdentity(), key, res, setErr)
		err = setErr
	case err != nil:
		log.Printf("Check Get Count Failed: %v", err)
	case val >= xrl.LimitNum:
		timeleft := xrl.client.TTL(ctx, key).Val()
		out.IsAllowed = false
		out.Timeleft = uint64(timeleft)
		out.CountRemaining = xrl.LimitNum - val
		log.Printf("%v(GRPC): <-> %v(Redis): Exceed The Limit Number(Reset:%v)", in.GetIdentity(), key, timeleft)
		//err = LimitExceedError
	default:
		timeleft := xrl.client.TTL(ctx, key).Val()
		res, setErr := xrl.client.Set(ctx, key, val+increase, timeleft).Result()
		out.IsAllowed = true
		out.Timeleft = uint64(timeleft)
		out.CountRemaining = xrl.LimitNum - val - increase
		log.Printf("%v(GRPC) <-> %v(Redis), Result:%v(Err:%v)", in.GetIdentity(), key, res, setErr)
		err = setErr
	}
	return key, out, err
}

func (xrl *XlimitRedisLayer) Get(ctx context.Context, in *xlimit.XLimitGetRequest, out *xlimit.XLimitGetReply) (*xlimit.XLimitGetReply, error) {
	iter := xrl.client.Scan(ctx, 0, xrl.Prefix+":"+in.GetIdentity(), 0).Iterator()
	for iter.Next(ctx) {
		key := iter.Val()
		val, err := xrl.client.Get(ctx, key).Uint64()
		timeleft := xrl.client.TTL(ctx, key).Val()
		switch {
		case err != nil:
			log.Printf("Check Get Count Failed: %v", err)
		case err == redis.Nil || val == 0:
		default:
			result := xlimit.XLimitCheckReply{Identity: key, CountRemaining: xrl.LimitNum - val, Timeleft: uint64(timeleft), IsAllowed: xrl.LimitNum <= val}
			out.Results = append(out.Results, &result)
		}
	}
	err := iter.Err()
	log.Printf("%v", out)
	return out, err
}

func (xrl *XlimitRedisLayer) GetKeysOnly(ctx context.Context, in *xlimit.XLimitGetRequest, out *xlimit.XLimitGetReply) (*xlimit.XLimitGetReply, error) {
	iter := xrl.client.Scan(ctx, 0, xrl.Prefix+":"+in.GetIdentity(), 0).Iterator()
	for iter.Next(ctx) {
		result := xlimit.XLimitCheckReply{Identity: iter.Val()}
		out.Results = append(out.Results, &result)
	}
	err := iter.Err()
	log.Printf("%v", out)
	return out, err
}

func NewXlimitRedisLayer(client *redis.Client, options *XlimitRedisLayerOptions) *XlimitRedisLayer {
	return &XlimitRedisLayer{
		client:                  client,
		XlimitRedisLayerOptions: options,
	}
}

func NewXlimitRedisLayerFromEnv() *XlimitRedisLayer {
	var redisLayerOptions XlimitRedisLayerOptions
	if err := envconfig.Process("REDIS_LAYER", &redisLayerOptions); err != nil {
		log.Fatalln(err)
	}
	log.Println(redisLayerOptions)

	return NewXlimitRedisLayer(
		redis.NewClient(&redis.Options{Addr: os.Getenv("REDIS_ADDR"), Password: os.Getenv("REDIS_PASSWORD"), DB: 0}),
		&redisLayerOptions,
	)
}
