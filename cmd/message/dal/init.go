package dal

import "tiktok-server/cmd/message/dal/db"

// Init init dal
func Init() {
	db.Init() // mysql
}
