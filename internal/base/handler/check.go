package handler

import (
	"net/http"
	"news-api/internal/base/app"
	"news-api/pkg/server"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (b BaseHTTPHandler) IsStaging() bool {
	return b.AppConfig.IsStaging()
}

func (b BaseHTTPHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "ok",
		"data":    "application running",
	})
}

func (b BaseHTTPHandler) Test(ctx *app.Context) *server.Response {

	logrus.Infoln(ctx.APIReqID)

	return &server.Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    nil,
	}
}
