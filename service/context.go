package service

import (
	"net/http"
	"strings"

	"github.com/fitzix/assassin/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

type AsnGin struct {
	C  *gin.Context
	D  *gorm.DB
	L  *zap.SugaredLogger
}

func NewAsnGin(c *gin.Context) *AsnGin {
	return &AsnGin{
		C:  c,
		D: db,
		L:  logger,
	}
}

func (a *AsnGin) Page(stmt interface{}, page models.PageReq, count *int, dest interface{}) error {
	// query := stmt.Column
	// stmt.Column = []interface{}{"COUNT (*)"}
	// if err := stmt.LoadOne(count); err != nil {
	// 	return err
	// }
	//
	// stmt.Column = query
	// stmt.Offset(uint64(page.PerPage * (page.Page - 1))).Limit(uint64(page.PerPage))
	// _, err := stmt.Load(dest)
	// return err
	return nil
}

func (a *AsnGin) GetToken() models.Token {
	token, _ := a.C.Get("token")
	return *token.(*models.Token)
}

func (a *AsnGin) IsAuth() bool {
	authorization := a.C.GetHeader("Authorization")
	if authorization == "" {
		return false
	}
	if _, err := ParseToken(strings.TrimPrefix(authorization, "Bearer ")); err != nil {
		return false
	}
	return true
}

// Response setting gin.JSON
func (a *AsnGin) Response(code int, data interface{}) {
	a.C.JSON(http.StatusOK, models.Response{
		Code: code,
		Msg:  AsnStatusText(code),
		Data: data,
	})
}

func (a *AsnGin) Success(data interface{}) {
	a.Response(0, data)
}

func (a *AsnGin) Fail(code int, err error) {
	a.L.Warnf("response err: %d %s", code, err)
	a.Response(code, nil)
}

func (a *AsnGin) SuccessWithPage(data interface{}) {
	a.Response(0, data)
}
