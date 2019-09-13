package github

import (
	"context"
	"io/ioutil"
	"mime/multipart"
	"path"
	"path/filepath"

	"github.com/fitzix/assassin/models"
	"github.com/fitzix/assassin/utils"
	"github.com/fitzix/assassin/utils/encrypt"
	"github.com/google/go-github/v28/github"
	"golang.org/x/oauth2"
)

type githubClient struct {
	client *github.Client
	models.Github
	ctx context.Context
}

var client *githubClient

func InitGithubClient() {
	conf := utils.GetConf().Github

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: conf.Token,
	})
	tc := oauth2.NewClient(ctx, ts)

	client = &githubClient{
		client: github.NewClient(tc),
		Github: conf,
		ctx:    ctx,
	}
}

func (c *githubClient) UploadToGithub(content []byte, ext string) (string, error) {
	fileName := encrypt.GetNanoId() + ext
	opts := &github.RepositoryContentFileOptions{
		Message: github.String("upload file by asins.xyz"),
		Content: content,
		Branch:  github.String(c.Branch),
		// Committer: &github.CommitAuthor{Name: github.String("FirstName LastName"), Email: github.String("user@example.com")},
	}
	_, _, err := c.client.Repositories.CreateFile(c.ctx, c.Owner, c.Repo, path.Join(c.FilePath, fileName), opts)
	if err != nil {
		return "", err
	}

	return fileName, nil
}

func (c *githubClient) UploadFromFileHeader(file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	ret, err := ioutil.ReadAll(src)
	if err != nil {
		return "", err
	}

	ext := filepath.Ext(file.Filename)

	return c.UploadToGithub(ret, ext)
}

func GetGithubClient() *githubClient {
	return client
}
