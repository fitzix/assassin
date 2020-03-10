package controllers

import (
	"github.com/fitzix/assassin/models"
	"github.com/fitzix/assassin/service"
	"github.com/fitzix/assassin/utils"
	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
	a := service.NewAsnGin(c)
	var up models.UserLoginReq
	if err := c.ShouldBind(&up); err != nil {
		a.Fail(service.StatusParamErr, err)
		return
	}
	var user models.User
	// if err := a.Db.Get(&user, `SELECT * FROM "user" WHERE "name" = $1`, up.UserName); err != nil {
	// 	a.Fail(service.StatusUserNotExist, err)
	// 	return
	// }

	if !utils.CheckPass(user.Password, up.Password) {
		a.Fail(service.StatusUserWrongPwd, nil)
		return
	}

	token, err := service.GenJwt(models.Token{
		Uid:  user.UID,
		Code: user.Code,
	})
	if err != nil {
		a.Fail(service.StatusBadRequest, err)
		return
	}

	a.Success(models.UserLoginRsp{
		User:  user,
		Token: token,
	})
}

func UserCreate(c *gin.Context) {
	a := service.NewAsnGin(c)
	var up models.UserLoginReq
	if err := c.ShouldBind(&up); err != nil {
		a.Fail(service.StatusParamErr, err)
		return
	}
	encPwd, err := utils.EncryptPass(up.Password)
	if err != nil {
		a.Fail(service.StatusBadRequest, err)
		return
	}

	user := models.User{
		UID:      utils.GenNanoId(),
		Name:     up.UserName,
		Password: string(encPwd),
		RoleId:   1,
	}
	// if _, err := a.Db.Exec(
	// 	`INSERT INTO "user" ("uid","name","password","role_id") VALUES ($1,$2,$3,$4)`,
	// 	user.UID,
	// 	user.Name,
	// 	user.Password,
	// 	user.RoleId,
	// ); err != nil {
	// 	a.Fail(service.StatusBadRequest, err)
	// 	return
	// }
	a.Success(user)
}
