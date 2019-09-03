package models

type Config struct {
	Db
}

func NewConfig() *Config {
	return &Config{Db{
		Host:     "127.0.0.1",
		Port:     5432,
		User:     "postgres",
		Password: "131833",
		Dbname:   "assassin",
	}}
}

type Db struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}


