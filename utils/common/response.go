package common

import (
	"gBlog/global"
	"gBlog/utils/error_code"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 封装页面响应

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

const (
	Success = 0
	ERROR   = 7
)

func Result(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func Ok(c *gin.Context) {
	Result(Success, map[string]any{}, "success", c)
}

func OkWithDetail(data any, msg string, c *gin.Context) {
	Result(Success, data, msg, c)
}

func OkWithData(data any, c *gin.Context) {
	Result(Success, data, "success", c)
}

func OkWithMessage(msg string, c *gin.Context) {
	Result(Success, map[string]any{}, msg, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, map[string]any{}, "fail", c)
}

func FailWithMessage(msg string, c *gin.Context) {
	Result(ERROR, map[string]any{}, msg, c)
}

func FailWithCode(code error_code.ErrorCode, c *gin.Context) {
	msg, ok := global.ErrorMap[code]
	if ok {
		Result(int(code), map[string]any{}, msg, c)
	}
	Result(ERROR, map[string]any{}, "unknown error", c)
}
