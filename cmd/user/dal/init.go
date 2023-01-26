package dal

import "tiktok-server/cmd/user/dal/db"

// Init init dal
func Init() {
	db.Init() // mysql
}
