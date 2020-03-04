package models

type Config struct {
	Salt string
	Env  string
	Db
	Github
	Jwt
	Encrypt
}

type Db struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}

type Github struct {
	GithubServer string
	Owner        string
	Repo         string
	Branch       string
	ImgPath      string
	ArticlePath  string
	AppDescPath  string
	Token        string
}

type Jwt struct {
	Issuer  string
	Expires int // hour
	Secret  string
}

type Encrypt struct {
	Key string
}
