package dal

import "tiktok-server/cmd/feed/dal/db"

// Init init dal
func Init() {
	db.Init() // mysql
}
