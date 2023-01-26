package dal

import "tiktok-server/cmd/publish/dal/db"

// Init init dal
func Init() {
	db.Init() // mysql
}
