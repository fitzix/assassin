package models

type Config struct {
	Salt string
	Db
	Github
	Jwt
	Encrypt
}

func NewConfig() *Config {
	return &Config{
		Salt: "404 page not found",
		Db: Db{
			Host:     "127.0.0.1",
			Port:     5432,
			User:     "postgres",
			Password: "131833",
			Dbname:   "assassin",
		},
		Github: Github{
			GithubServer: "https://raw.githubusercontent.com",
			Owner:        "asins-xyz",
			Repo:         "assassin",
			Token:        "e89ebc4f50935201c88349726c6709f8a2e09e85",
			Branch:       "master",
			ImgPath:      "post/img",
			ArticlePath:  "post/article",
			AppDescPath:  "post/app",
		},
		Jwt: Jwt{
			Issuer:  "asn.xyz",
			Expires: 24 * 3,
			Secret:  "asn.io",
		},
		Encrypt: Encrypt{
			Key: "3C221351CA73FFA6",
		},
	}
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
	Expires int
	// hour
	Secret string
}

type Encrypt struct {
	Key string
}
