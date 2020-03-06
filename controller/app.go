package controller

import (
	"context"

	"github.com/fitzix/assassin/ent"
	"github.com/fitzix/assassin/ent/app"
	"github.com/fitzix/assassin/ent/hot"
	"github.com/fitzix/assassin/ent/predicate"
	"github.com/fitzix/assassin/models"
	"github.com/fitzix/assassin/service"
	"github.com/labstack/echo/v4"
)

func AppList(e echo.Context) error {
	c := e.(*service.Context)
	up := models.AppListReq{
		PageReq: models.NewPageReq(),
	}
	if err := c.Bind(&up); err != nil {
		return c.Err(service.StatusWebParamErr, err)
	}
	var (
		condition []predicate.App
		// order     ent.Order
	)
	if up.Name != "" {
		condition = append(condition, app.NameContains(up.Name))
	}

	if err := app.TypeValidator(up.Type); err == nil {
		condition = append(condition, app.TypeEQ(up.Type))
	}

	// appQuery := service.GetDB().App.Query().Where(condition...)
	// var down []*ent.App

	if up.Order == "hot" {
		c.Logger().Info(len(condition))
		r, _:=service.GetDB().Hot.Query().WithApp(func(q *ent.AppQuery) {
			q.Where(condition...)
		}).Order(ent.Desc(hot.FieldHot)).Limit(2).All(context.Background())
		// order = ent.Desc(hot.FieldHot)
		return c.Success(r)
	} else {
		// order = ent.Desc(version.FieldCreatedAt)
	}
	return nil

	// ctx := context.Background()
	// total, err := appQuery.Count(ctx)
	// if err != nil {
	// 	return c.Err(service.StatusWebBadRequest, nil)
	// }
	//
	// down, err = appQuery.Order(order).Offset(up.PageSize * (up.PageNum - 1)).Limit(up.PageSize).All(ctx)
	// if err != nil {
	// 	return c.Err(service.StatusWebBadRequest, nil)
	// }
	// return c.SuccessWithPage(total, down)
}
