package service

import (
	"context"

	"github.com/fitzix/assassin/models"
	"github.com/google/go-github/v28/github"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

var (
	appConf      models.Config
	zapLogger    *zap.Logger
	githubClient *GithubClient
	dbInstance   *gorm.DB
)

func InitProject() {
	initConf()
	initGithubClient()
}

func initGithubClient() {
	conf := appConf.Github

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: conf.Token,
	})
	tc := oauth2.NewClient(ctx, ts)

	githubClient = &GithubClient{
		client: github.NewClient(tc),
		Github: conf,
		ctx:    ctx,
	}
}

func GetGithubClient() *GithubClient {
	return githubClient
}