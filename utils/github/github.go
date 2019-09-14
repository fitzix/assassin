package github

import (
	"context"
	"io/ioutil"
	"mime/multipart"
	"path"
	"path/filepath"

	"github.com/fitzix/assassin/models"
	"github.com/fitzix/assassin/service"
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

func (c *githubClient) UploadToGithub(fileName, filePath string, content []byte) (string, error) {
	opts := &github.RepositoryContentFileOptions{
		Message: github.String("upload file by asins.xyz"),
		Content: content,
		Branch:  github.String(c.Branch),
		// Committer: &github.CommitAuthor{Name: github.String("FirstName LastName"), Email: github.String("user@example.com")},
	}
	_, _, err := c.client.Repositories.CreateFile(c.ctx, c.Owner, c.Repo, path.Join(filePath, fileName), opts)
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

	fileByte, err := ioutil.ReadAll(src)
	if err != nil {
		return "", err
	}

	fileName := encrypt.GetNanoId() + filepath.Ext(file.Filename)

	return c.UploadToGithub(fileName, c.ImgPath, fileByte)
}

func (c *githubClient) CreateMdFile(fileName string, uploadType int) (string, error) {
	filePath := c.AppDescPath
	if uploadType == service.AsnUploadTypeArticle {
		filePath = c.ArticlePath
	}
	fileByte := service.GetTmplContent(uploadType)

	name, err := c.UploadToGithub(fileName+".md", filePath, fileByte)

	if err != nil {
		utils.GetLogger().Sugar().Warnf("upload md file to github err: %s", err)
	}

	return name, err
}

func GetGithubClient() *githubClient {
	return client
}
