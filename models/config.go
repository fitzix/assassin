package models

type Config struct {
	Salt string
	Db
	Github
	Jwt
	Encrypt
	S3
}

type Db struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
	SSL      bool
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

type S3 struct {
	Endpoint        string
	ImgPrefix       string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool
	Bucket          string
}
