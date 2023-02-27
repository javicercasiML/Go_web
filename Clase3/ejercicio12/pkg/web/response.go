package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Data interface{} `json:"data"`
}

func Success(ctx *gin.Context, code int, data any) {
	ctx.JSON(code, Response{Data: data})
}

func Failed(ctx *gin.Context, status int, err error) {
	ctx.JSON(status, errorResponse{
		Status:  status,
		Code:    http.StatusText(status),
		Message: err.Error(),
	})
}
