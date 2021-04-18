package work

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

var rdb *redis.Client
var ctx = context.Background()

const (
	IpLimitDuration  = time.Minute
	IpLimitThreshold = 30
)

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func IsIpLimit(ip string) bool {
	val, err := rdb.Get(ctx, ip).Int()
	if err != nil || val < 0 {
		val = 0
	}
	if val >= IpLimitThreshold {
		return true
	}
	val++
	ttl, err := rdb.TTL(ctx, ip).Result()
	if ttl <= 0 {
		ttl = IpLimitDuration
		rdb.Set(ctx, ip, "0", IpLimitDuration)
	} else {
		rdb.Set(ctx, ip, string(val), -1)

	}
	return false
}
