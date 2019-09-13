package models

type Config struct {
	Db
	Github
}

func NewConfig() *Config {
	return &Config{
		Db{
			Host:     "127.0.0.1",
			Port:     5432,
			User:     "postgres",
			Password: "131833",
			Dbname:   "assassin",
		},
		Github{
			Owner:      "asins-xyz",
			Repo:       "assassin",
			Token:      "e89ebc4f50935201c88349726c6709f8a2e09e85",
			Branch:     "master",
			FilePath:   "post/img",
			GithubPath: "https://raw.githubusercontent.com",
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
