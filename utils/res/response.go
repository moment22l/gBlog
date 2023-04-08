package res

import (
	"gBlog/global"
	"gBlog/utils"
	"gBlog/utils/error_code"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 封装页面响应

// Response 响应
type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

// List 列表返回类型
type List struct {
	Count int64 `json:"count"`
	List  any   `json:"list"`
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

func OkWithList(list any, count int64, c *gin.Context) {
	OkWithData(List{
		Count: count,
		List:  list,
	}, c)
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

func FailWithError(err error, obj any, c *gin.Context) {
	msg := utils.GetValidMsg(err, obj)
	FailWithMessage(msg, c)
}
