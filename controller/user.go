package controller

import (
	"context"

	"github.com/fitzix/assassin/ent/user"
	"github.com/fitzix/assassin/models"
	"github.com/fitzix/assassin/service"
	"github.com/fitzix/assassin/utils"
	"github.com/labstack/echo/v4"
)

func UserLogin(e echo.Context) error {
	c := e.(*service.Context)
	var up models.UserLoginReq
	if err := c.ShouldBind(&up); err != nil {
		return c.Err(service.StatusWebParamErr, err)
	}
	encPwd, err := utils.EncryptPass(up.Password)
	if err != nil {
		return c.Err(service.StatusWebBadRequest, err)
	}
	u, err := service.GetDB().User.Query().Where(user.Name(up.UserName), user.Password(string(encPwd))).Only(context.Background())
	if err != nil {
		return c.Err(service.StatusWebAuthWrongPwd, err)
	}
	token := models.Token{
		Uid:  u.ID,
		Code: u.Code,
	}
	t, err := service.GenJwt(token)
	if err != nil {
		return c.Err(service.StatusWebBadRequest, err)
	}
	return c.Success(models.UserLoginRsp{Token: t})
}

func UserCreate(e echo.Context) error {
	c := e.(*service.Context)
	var up models.UserLoginReq
	if err := c.ShouldBind(&up); err != nil {
		return c.Err(service.StatusWebParamErr, err)
	}

	encPass, err := utils.EncryptPass(up.Password)
	if err != nil {
		return c.Err(service.StatusWebBadRequest, err)
	}
	u, err := service.GetDB().User.Create().SetName(up.UserName).SetPassword(string(encPass)).SetRoleID(1).Save(context.Background())
	if err != nil {
		return c.Err(service.StatusWebBadRequest, err)
	}
	return c.Success(u)
}
