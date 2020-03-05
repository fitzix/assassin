package controller

import (
	"github.com/labstack/echo/v4"
)

func AppList(e echo.Context) error {
	// c := e.(*service.Context)
	// up := models.AppListReq{
	// 	PageReq: models.NewPageReq(),
	// }
	// if err := c.Bind(&up); err != nil {
	// 	return c.Err(service.StatusWebParamErr, err)
	// }
	// var (
	// 	condition []predicate.App
	// 	order     ent.Order
	// )
	// if up.Name != "" {
	// 	condition = append(condition, app.NameContains(up.Name))
	// }
	//
	// if err := app.TypeValidator(up.Type); err == nil {
	// 	condition = append(condition, app.TypeEQ(up.Type))
	// }
	//
	// if up.Order == "hot" {
	// 	order = ent.Desc(hot.FieldHot)
	// } else {
	// 	order = ent.Desc(version.FieldCreatedAt)
	// }
	//
	// service.GetDB().Hot.Query().All()
	//
	// appQuery := service.GetDB().App.Query().Where(condition...)
	// ctx := context.Background()
	// total, err := appQuery.Count(ctx)
	// if err != nil {
	// 	return c.Err(service.StatusWebBadRequest, nil)
	// }
	// down, err := appQuery.Offset(up.PageSize * (up.PageNum - 1)).Limit(up.PageSize).Order(order).All(ctx)
	// if err != nil {
	// 	return c.Err(service.StatusWebBadRequest, nil)
	// }
	// return c.SuccessWithPage(total, down)
	return nil
}
