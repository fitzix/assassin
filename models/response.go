package models

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type PageDown struct {
	Total int         `json:"total"`
	Info  interface{} `json:"info"`
}
