package controller

import (
	"github.com/labstack/echo/v4"
)

func UserLogin(e echo.Context) error {
	// c := e.(*service.Context)
	// var up models.UserLoginReq
	// if err := c.ShouldBind(&up); err != nil {
	// 	return c.Err(service.StatusWebParamErr, err)
	// }
	// u, err := service.GetDB().User.Query().WithRole().Where(user.Name(up.UserName)).Only(context.Background())
	// if err != nil {
	// 	return c.Err(service.StatusWebAuthWrongPwd, err)
	// }
	// if !utils.CheckPass(u.Password, up.Password) {
	// 	return c.Err(service.StatusWebBadRequest, err)
	// }
	// t, err := service.GenJwt(models.Token{
	// 	Uid:  u.ID,
	// 	Code: u.Code,
	// })
	// if err != nil {
	// 	return c.Err(service.StatusWebBadRequest, err)
	// }
	// return c.Success(models.UserLoginRsp{
	// 	User:  u,
	// 	Token: t,
	// })
	return nil
}

func UserCreate(e echo.Context) error {
	// c := e.(*service.Context)
	// var up models.UserLoginReq
	// if err := c.ShouldBind(&up); err != nil {
	// 	return c.Err(service.StatusWebParamErr, err)
	// }
	//
	// encPass, err := utils.EncryptPass(up.Password)
	// if err != nil {
	// 	return c.Err(service.StatusWebBadRequest, err)
	// }
	// u, err := service.GetDB().User.Create().SetName(up.UserName).SetPassword(string(encPass)).SetRoleID(1).Save(context.Background())
	// if err != nil {
	// 	return c.Err(service.StatusWebBadRequest, err)
	// }
	// t, err := service.GenJwt(models.Token{
	// 	Uid:  u.ID,
	// 	Code: u.Code,
	// })
	// return c.Success(models.UserLoginRsp{
	// 	User:  u,
	// 	Token: t,
	// })
	return nil
}
