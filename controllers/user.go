package controllers

import (
	"github.com/fitzix/assassin/models"
	"github.com/fitzix/assassin/service"
	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
	a := service.NewAsnGin(c)
	var up service.UserLogin
	if err := c.ShouldBind(&up); err != nil {
		a.Fail(service.StatusWebParamErr, err)
		return
	}
	var user models.User
	encPwd := service.PassEncrypt(up.Password)
	if err := a.D.Find(&user, "name = ? AND password = ?", up.UserName, encPwd).Error; err != nil {
		a.Fail(service.StatusWebAuthWrongPwd, err)
		return
	}

	token := service.Token{
		Uid:  user.ID,
		Code: user.Code,
	}
	tokenString, err := service.GenJwt(token)
	if err != nil {
		a.Fail(service.StatusWebBadRequest, err)
		return
	}

	a.Success(tokenString)
}

func UserCreate(c *gin.Context) {
	a := service.NewAsnGin(c)
	var up models.User
	if err := c.ShouldBind(&up); err != nil {
		a.Fail(service.StatusWebParamErr, err)
		return
	}
	up.Password = service.PassEncrypt(up.Password)
	if err := a.D.Create(&up).Error; err != nil {
		a.Fail(service.StatusWebBadRequest, err)
		return
	}
	a.Success(up)
}
