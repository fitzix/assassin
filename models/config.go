package models

type Config struct {
	Salt string
	Db
	Github
	Jwt
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
			Owner:      "asins-xyz",
			Repo:       "assassin",
			Token:      "e89ebc4f50935201c88349726c6709f8a2e09e85",
			Branch:     "master",
			FilePath:   "post/img",
			GithubPath: "https://raw.githubusercontent.com",
		},
		Jwt: Jwt{
			Issuer:  "asn.xyz",
			Expires: 24 * 7,
			Secret:  "asn.io",
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
	Owner      string
	Repo       string
	Token      string
	Branch     string
	FilePath   string
	GithubPath string
}

type Jwt struct {
	Issuer  string
	Expires int
	// hour
	Secret string
}
