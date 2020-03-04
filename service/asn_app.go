package service

import (
	"strconv"

	"github.com/fitzix/assassin/ent"
	"github.com/fitzix/assassin/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type Context struct {
	echo.Context
	Db *ent.Client
}

type AsnGin struct {
	C *gin.Context
	D *gorm.DB
	L *zap.SugaredLogger
}

func NewAsnGin(c *gin.Context) *AsnGin {
	return &AsnGin{
		C: c,
		D: dbInstance,
		L: zapLogger.Sugar(),
	}
}

func (c *Context) Resp(code int, data interface{}) error {
	return c.JSON(200, models.Response{
		Code: code,
		Msg:  AsnStatusText(code),
		Data: data,
	})
}

func (c *Context) Success(data interface{}) error {
	return c.Resp(StatusSuccess, data)
}

func (c *Context) SuccessWithPage(total int, data interface{}) error {
	return c.Success(models.PageDown{
		Total: total,
		Info:  data,
	})
}

func (c *Context) Err(code int, err error) error {
	c.Logger().Warnf("response err, code: %d, err: %s", code, err)
	return c.Resp(code, nil)
}

func (c *Context) GetToken() (models.Token, bool) {
	v, ok := c.Get("token").(models.Token)
	return v, ok
}

func (a *AsnGin) Page(query *gorm.DB, data interface{}, count interface{}) error {
	size, err := strconv.Atoi(a.C.Query("page_size"))
	if err != nil {
		size = 20
	}
	num, err := strconv.Atoi(a.C.Query("page_num"))
	if err != nil {
		num = 1
	}

	if err := query.Model(data).Count(count).Error; err != nil {
		return err
	}
	if err := query.Limit(size).Offset(size * (num - 1)).Find(data).Error; err != nil {
		return err
	}
	return nil
}

func (a *AsnGin) GetToken() Token {
	token, _ := a.C.Get("token")
	return *token.(*Token)
}

func (a *AsnGin) IsAuth() bool {
	_, exists := a.C.Get("token")
	return exists
}
