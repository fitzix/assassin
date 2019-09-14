package controllers

import (
	"github.com/fitzix/assassin/models"
	"github.com/fitzix/assassin/service"
	"github.com/fitzix/assassin/service/model"
	"github.com/fitzix/assassin/utils/encrypt"
	"github.com/gin-gonic/gin"
)

// func (r *mutationResolver) Login(ctx context.Context, input models.Login) (*models.LoginResp, error) {
// 	panic("implement me")
// }

func UserLogin(c *gin.Context) {
	a := service.NewAsnGin(c)
	var up model.UserLogin
	if err := c.ShouldBind(&up); err != nil {
		a.Fail(service.StatusWebParamErr, err)
		return
	}
	var user models.User
	encPwd := encrypt.PassEncrypt(up.Password)
	if err := a.D.Find(&user, "name = ? AND password = ?", up.UserName, encPwd).Error; err != nil {
		a.Fail(service.StatusWebAuthWrongPwd, err)
		return
	}

	token := model.Token{
		Uid:  user.ID,
		Code: user.Code,
	}
	tokenString, err := encrypt.GenJwt(token)
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
	up.Password = encrypt.PassEncrypt(up.Password)
	if err := a.D.Create(&up).Error; err != nil {
		a.Fail(service.StatusWebBadRequest, err)
		return
	}
	a.Success(up)
}
