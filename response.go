package main

type Response struct {
	Code    string      `json:"code"`
	Method  string      `json:"method"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
