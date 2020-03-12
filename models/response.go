package models

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type PageReq struct {
	Page    int `json:"page" form:"page,default=1"`
	PerPage int `json:"perPage" form:"perPage,default=20"`
}

type PageRsp struct {
	Total int       `json:"total"`
	Info  interface{} `json:"info"`
}
