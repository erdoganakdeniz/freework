package model

type Response struct {
	Code    int         `json:"code"`
	Method  string      `json:"method"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
