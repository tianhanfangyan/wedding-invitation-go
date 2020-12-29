package models

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 返回值
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// 分页值
type Page struct {
	PageNumber int64       `json:"page_number"`
	PageSize   int64       `json:"page_size"`
	TotalPage  int64       `json:"total_page"`
	TotalCount int64       `json:"total_count"`
	Data       interface{} `json:"data"`
}

const (
	SUCCESS = 0
	ERROR   = -1
)

func Result(code int, message string, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		message,
		data,
	})
}

func RetSuccess(c *gin.Context) {
	Result(SUCCESS, "OK", map[string]interface{}{}, c)
}

func RetSuccessWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, message, map[string]interface{}{}, c)
}

func RetSuccessWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, "OK", data, c)
}

func RetSuccessWithDetail(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, message, data, c)
}

func RetFail(c *gin.Context) {
	Result(ERROR, "FAIL", map[string]interface{}{}, c)
}

func RetFailWithMessage(message string, c *gin.Context) {
	Result(ERROR, message, map[string]interface{}{}, c)
}

func RetFailWithData(data interface{}, c *gin.Context) {
	Result(ERROR, "FAIL", data, c)
}

func RetFailWithDetail(data interface{}, message string, c *gin.Context) {
	Result(ERROR, message, data, c)
}
