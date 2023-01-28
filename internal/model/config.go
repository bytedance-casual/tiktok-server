package model

type Config struct {
	Server  *Server
	Database *Database
	Redis    *Redis
}

type Server struct {
	Address string
}

type Database struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string `yaml:"db-name"`
	Settings string
}

type Redis struct {
	Host     string
	Port     int
	Password string
	DBIndex  int `yaml:"db-index"`
}
