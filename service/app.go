package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AsnGin struct {
	c *gin.Context
}

type Response struct {
	Code AsnStatusCode `json:"code"`
	Msg  string        `json:"msg"`
	Data interface{}   `json:"data"`
}

// Response setting gin.JSON
func (a *AsnGin) Response(code AsnStatusCode, data interface{}) {
	a.c.JSON(http.StatusOK, Response{
		Code: code,
		Msg: AsnStatusText(code),
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
