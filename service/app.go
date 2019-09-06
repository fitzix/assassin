package service

import (
	"net/http"
	"strconv"

	"github.com/fitzix/assassin/db"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type AsnGin struct {
	C *gin.Context
	D *gorm.DB
}

func NewAsnGin(c *gin.Context) *AsnGin {
	return &AsnGin{C: c, D: db.GetDB() }
}

type Response struct {
	Code AsnStatusCode `json:"code"`
	Msg  string        `json:"msg"`
	Data interface{}   `json:"data"`
}

// Response setting gin.JSON
func (a *AsnGin) Response(code AsnStatusCode, data interface{}) {
	a.C.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  AsnStatusText(code),
		Data: data,
	})
	return
}

func (a *AsnGin) Success(data interface{}) {
	a.Response(0, data)
	return
}

func (a *AsnGin) Fail(code AsnStatusCode) {
	a.Response(code, nil)
	return
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
