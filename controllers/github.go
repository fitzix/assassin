package controllers

import (
	"path"

	"github.com/fitzix/assassin/service"
	"github.com/fitzix/assassin/utils"
	"github.com/fitzix/assassin/utils/github"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {

}

func Upload(c *gin.Context) {
	a := service.NewAsnGin(c)
	file, err := c.FormFile("file")
	if err != nil {
		a.Fail(service.StatusWebBadRequest, err)
		return
	}

	fileName, err := github.GetGithubClient().UploadFromFileHeader(file)

	if err != nil {
		a.Fail(service.StatusWebBadRequest, err)
		return
	}

	conf := utils.GetConf().Github
	a.Success(path.Join(conf.GithubPath, conf.FilePath, fileName))
}
