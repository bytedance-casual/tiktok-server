package redis

import (
	"github.com/go-redis/redis"
	"strconv"
	"tiktok-server/internal/conf"
)

var DB *redis.Client

func Init() {
	r := conf.Config.Redis
	DB = redis.NewClient(&redis.Options{
		Addr:     r.Host + ":" + strconv.Itoa(r.Port),
		Password: r.Password, // no password set
		DB:       r.DBIndex,  // use default DB
	})
}
