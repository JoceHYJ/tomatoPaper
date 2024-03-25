package util

import (
	"net/http"
	"tomatoPaper/web"
)

type GormResponse struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	Data    any    `json:"data"`
}

var gormResponse GormResponse

func HandleResponse(c *web.Context, code int, msg string, data any) {
	gormResponse.Code = code
	gormResponse.Message = msg
	gormResponse.Data = data
	c.RespJSON(http.StatusOK, gormResponse)
}
