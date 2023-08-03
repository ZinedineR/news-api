package handler

import (
	"fmt"
	"net/http"
	"time"

	"news-api/pkg/jwthelper"

	"github.com/sirupsen/logrus"

	"news-api/app/appconf"
	"news-api/internal/base/app"
	baseModel "news-api/pkg/db"
	"news-api/pkg/server"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HandlerFn func(ctx *app.Context) *server.Response
type HandlerFnInterface func(ctx *app.Context) *server.ResponseInterface

type BaseHTTPHandler struct {
	Handlers  interface{}
	DB        *gorm.DB
	AppConfig *appconf.Config
	BaseModel *baseModel.PostgreSQLClientRepository
}

func NewBaseHTTPHandler(db *gorm.DB,
	appConfig *appconf.Config,
	baseModel *baseModel.PostgreSQLClientRepository,
) *BaseHTTPHandler {
	return &BaseHTTPHandler{
		DB:        db,
		AppConfig: appConfig,
		BaseModel: baseModel,
	}
}

// AsJson to response custom message: 200, 201 with message (Mobile use 500 error)
func (b BaseHTTPHandler) AsJson(ctx *app.Context, status int, message string, data interface{}) *server.Response {

	return &server.Response{
		Status:       status,
		Message:      message,
		Data:         data,
		ResponseType: server.DefaultResponseType,
	}
}

func (b BaseHTTPHandler) AsJsonInterface(ctx *app.Context, status int, data interface{}) *server.ResponseInterface {

	return &server.ResponseInterface{
		Status: status,
		Data:   data,
	}
}

// ThrowExceptionJson for some exception not handle in Yii2 framework
func (b BaseHTTPHandler) ThrowExceptionJson(ctx *app.Context, status, code int, name, message string) *server.Response {
	return &server.Response{
		Status:  status,
		Message: "",
		Log:     nil,
	}
}

func (b BaseHTTPHandler) MoodleAuthentication(c *gin.Context) (*app.Context, error) {
	return app.NewContext(c, b.AppConfig), nil
}

func (b BaseHTTPHandler) MoodleRunAction(handler HandlerFnInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		ctx, err := b.MoodleAuthentication(c)
		if err != nil {
			logrus.Errorln(fmt.Sprintf("REQUEST ID: %s , message: Unauthorized", ctx.APIReqID))
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "Unauthorized",
				"data":    err.Error(),
			})
			return
		}
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			logrus.Errorln(fmt.Sprintf("REQUEST ID: %s , message: Unauthorized", ctx.APIReqID))
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "request does not contain an access token",
			})
			return
		}
		message, err := jwthelper.ValidateToken(tokenString)
		if err != nil {
			logrus.Errorln(fmt.Sprintf("REQUEST ID: %s , message: Unauthorized", ctx.APIReqID))
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": message,
				"data":    err,
			})
			return
		}
		defer func() {
			if err0 := recover(); err0 != nil {
				logrus.Errorln(err0)
				c.JSON(http.StatusInternalServerError, gin.H{
					"status":  http.StatusInternalServerError,
					"message": "Request is halted unexpectedly, please contact the administrator.",
					"data":    nil,
				})
			}
		}()

		resp := handler(ctx)
		httpStatus := resp.Status

		if resp.Data == nil {
			c.Status(httpStatus)
			return
		}
		end := time.Now().Sub(start)
		logrus.Infoln(fmt.Sprintf("REQUEST ID: %s , LATENCY: %vms", ctx.APIReqID, end.Milliseconds()))
		c.JSON(httpStatus, resp.Data)

	}
}
