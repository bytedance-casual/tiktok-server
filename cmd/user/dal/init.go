package dal

import (
	"tiktok-server/cmd/user/dal/db"
	"tiktok-server/cmd/user/dal/redis"
)

// Init init dal
func Init() {
	db.Init() // mysql
	redis.Init()
}
