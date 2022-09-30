package response

import (
	"awesomeProject/serializer"
	"github.com/gin-gonic/gin"
	"net/http"
)

const MessageOk = "请求成功"

const MessageNoDBRecords = "没有查询到数据"

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type PageResult[T serializer.BaseParam] struct {
	*serializer.PageParam
	Records []T `json:"records"`
	Count   int `json:"count"`
}

type ListResult[T serializer.BaseParam] struct {
	Records []T `json:"records"`
	Count   int `json:"count"`
}

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(http.StatusOK, nil, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(http.StatusOK, nil, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(http.StatusOK, data, "操作成功", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(http.StatusOK, data, message, c)
}

func Fail(c *gin.Context) {
	Result(http.StatusInternalServerError, nil, "操作失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(http.StatusInternalServerError, nil, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(http.StatusInternalServerError, data, message, c)
}
