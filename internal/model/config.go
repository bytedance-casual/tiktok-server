package model

type Config struct {
	Server   Server
	Database Database
}

type Server struct {
	Address string
}

type Database struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
	Settings string
}
