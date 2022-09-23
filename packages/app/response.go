package app

import (
	"net/http"

	"github.com/201R/go_api_boilerplate/packages/logger"
	"github.com/gin-gonic/gin"
)

// Response object as HTTP response
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// HTTPRes normalize HTTP Response format
func HTTPRes(c *gin.Context, httpCode int, msg string, data interface{}) {

	if httpCode == http.StatusInternalServerError {
		logger.Error(msg)
		msg = "Internal server error "
	}

	c.JSON(httpCode, Response{
		Code: httpCode,
		Msg:  msg,
		Data: data,
	})

}
