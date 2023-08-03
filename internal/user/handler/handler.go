package handler

import (
	"net/http"
	"time"

	"news-api/internal/user/domain"
	UserService "news-api/internal/user/service"

	"news-api/internal/base/app"
	"news-api/internal/base/handler"
	"news-api/pkg/jwthelper"
	"news-api/pkg/mail"
	"news-api/pkg/server"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type HTTPHandler struct {
	App         *handler.BaseHTTPHandler
	UserService UserService.Service
}

func NewHTTPHandler(handler *handler.BaseHTTPHandler, UserService UserService.Service) *HTTPHandler {
	return &HTTPHandler{
		App:         handler,
		UserService: UserService,
	}
}

// Handler Basic Method ======================================================================================================

func (h HTTPHandler) AsErrorDefault(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"message": message,
	})
}

func (h HTTPHandler) AsInvalidClientIdError(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"responseCode":    "400",
		"responseMessage": "invalid clientid",
	})
}

func (h HTTPHandler) AsInvalidClientIdAccessTokenError(ctx *gin.Context) {
	ctx.JSON(http.StatusUnauthorized, gin.H{
		"responseCode":    "4010000",
		"responseMessage": "Invalid Client Key",
	})
}

func (h HTTPHandler) AsInvalidPrivateKeyError(ctx *gin.Context) {
	ctx.JSON(http.StatusUnauthorized, gin.H{
		"responseCode":    "4010000",
		"responseMessage": "Invalid Private Key",
	})
}

func (h HTTPHandler) AsDatabaseError(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"responseCode":    "500",
		"responseMessage": "Error in database",
	})
}

func (h HTTPHandler) AsNotVerfied(ctx *gin.Context) {
	ctx.JSON(http.StatusUnauthorized, gin.H{
		"responseCode":    "401",
		"responseMessage": "Account still not verified",
	})
}

func (h HTTPHandler) AsDuplicateEmail(ctx *gin.Context) {
	ctx.JSON(http.StatusUnauthorized, gin.H{
		"responseCode":    "401",
		"responseMessage": "Another account with same email/username already created",
	})
}

func (h HTTPHandler) AsDataNotFound(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, gin.H{
		"responseCode":    "404",
		"responseMessage": "Data not Found",
	})
}

func (h HTTPHandler) AsJWTExist(ctx *gin.Context) {
	ctx.JSON(http.StatusUnauthorized, gin.H{
		"responseCode":    "401",
		"responseMessage": "You already login before",
	})
}

func (h HTTPHandler) AsPasswordUnmatched(ctx *gin.Context) {
	ctx.JSON(http.StatusUnauthorized, gin.H{
		"responseCode":    "401",
		"responseMessage": "Password Unmatched",
	})
}

func (h HTTPHandler) AsHashError(ctx *gin.Context) {
	ctx.JSON(http.StatusUnauthorized, gin.H{
		"responseCode":    "500",
		"responseMessage": "Error in hashing",
	})
}

func (h HTTPHandler) AsDataUnauthorized(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusUnauthorized, gin.H{
		"responseCode":    "401",
		"responseMessage": message,
	})
}

func (h HTTPHandler) AsEmailNotFound(ctx *gin.Context) {
	ctx.JSON(http.StatusUnauthorized, gin.H{
		"responseCode":    "401",
		"responseMessage": "Can't send email, contact admin for verification",
	})
}

func (h HTTPHandler) AsInvalidPublicKeyError(ctx *gin.Context) {
	ctx.JSON(http.StatusUnauthorized, gin.H{
		"responseCode":    "4010000",
		"responseMessage": "Invalid Public Key",
	})
}

func (h HTTPHandler) AsInvalidSignatureError(ctx *gin.Context) {
	ctx.JSON(http.StatusUnauthorized, gin.H{
		"responseCode":    "4017300",
		"responseMessage": "Invalid Token (B2B)",
	})
}

func (h HTTPHandler) AsRequiredTimeStampError(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"responseCode":    "4000000",
		"responseMessage": "The timestamp field is required.",
	})
}

func (h HTTPHandler) AsInvalidFieldTimeStampError(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"responseCode":    "4000000",
		"responseMessage": "Invalid Field Format Timestamp",
	})
}

func (h HTTPHandler) AsInvalidLengthTimeStampError(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"responseCode":    "4000000",
		"responseMessage": "The field timestamp must be a string or array type with a maximum length of '25'.",
	})
}

func (h HTTPHandler) AsInvalidClientSecretError(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"responseCode":    "4010000",
		"responseMessage": "Invalid Client Secret",
	})
}

func (h HTTPHandler) AsInvalidHttpMethodError(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"responseCode":    "4010000",
		"responseMessage": "http methods is invalid",
	})
}

func (h HTTPHandler) AsInvalidJsonFormat(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"responseCode":    "400",
		"responseMessage": msg,
	})
}

func (h HTTPHandler) AsRequiredClientSecretError(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"responseCode":    "4000000",
		"responseMessage": "The clientSecret field is required.",
	})
}

func (h HTTPHandler) AsRequiredClientIdError(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"responseCode":    "4000000",
		"responseMessage": "The param ID is required.",
	})
}

func (h HTTPHandler) AsRequiredGrantTypeError(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"responseCode":    "4007302",
		"responseMessage": "Bad Request. The grantType field is required.",
	})
}

func (h HTTPHandler) AsRequiredGrantTypeClientCredentialsError(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"responseCode":    "4007300",
		"responseMessage": "grant_type must be set to client_credentials",
	})
}

func (h HTTPHandler) AsRequiredSignatureError(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"responseCode":    "4000000",
		"responseMessage": "The signature field is required.",
	})
}

func (h HTTPHandler) AsRequiredPrivateKeyError(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"responseCode":    "4000000",
		"responseMessage": "The privateKey field is required.",
	})
}

func (h HTTPHandler) AsRequiredContentTypeError(ctx *gin.Context) {
	ctx.JSON(http.StatusUnsupportedMediaType, gin.H{
		"responseCode":    "4000000",
		"responseMessage": "Content Type application/json is required.",
	})
}

func (h HTTPHandler) AsInvalidTokenError(ctx *gin.Context) {
	ctx.JSON(http.StatusUnauthorized, gin.H{
		"responseCode":    "4010001",
		"responseMessage": "Access Token Invalid",
	})
}

func (h HTTPHandler) AsRequiredBearer(ctx *gin.Context) {
	ctx.JSON(http.StatusUnauthorized, gin.H{
		"responseCode":    "4000002",
		"responseMessage": "Bearer authorization is required",
	})
}

func (h HTTPHandler) AsRequiredHttpMethodError(ctx *gin.Context) {
	ctx.JSON(http.StatusUnsupportedMediaType, gin.H{
		"responseCode":    "4000000",
		"responseMessage": "The HttpMethod field is required.",
	})
}

func (h HTTPHandler) AsRequiredEndpoinUrlError(ctx *gin.Context) {
	ctx.JSON(http.StatusUnsupportedMediaType, gin.H{
		"responseCode":    "4000000",
		"responseMessage": "The EndpointUrl field is required.",
	})
}

func (h HTTPHandler) AsRequiredAccessTokenError(ctx *gin.Context) {
	ctx.JSON(http.StatusUnsupportedMediaType, gin.H{
		"responseCode":    "4000000",
		"responseMessage": "The AccessToken field is required.",
	})
}
func (h HTTPHandler) AsRequiredBodyError(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"responseCode":    "4000000",
		"responseMessage": "A non-empty request body is required.",
	})
}

// Data Not Found return AsJsonInterface 404 when data doesn't exist
func (h HTTPHandler) DataNotFound(ctx *app.Context) *server.ResponseInterface {
	type Response struct {
		StatusCode int         `json:"responseCode"`
		Message    interface{} `json:"responseMessage"`
	}
	resp := Response{
		StatusCode: http.StatusNotFound,
		Message:    "Data not found in database.",
	}
	return h.App.AsJsonInterface(ctx, http.StatusNotFound, resp)

}

// DataReadError return AsJsonInterface error if persist a problem in declared condition
func (h HTTPHandler) DataReadError(ctx *app.Context, code int, description string) *server.ResponseInterface {
	type Response struct {
		StatusCode int         `json:"responseCode"`
		Message    interface{} `json:"responseMessage"`
	}
	resp := Response{
		StatusCode: code,
		Message:    description,
	}
	return h.App.AsJsonInterface(ctx, code, resp)
}

// AsJson always return httpStatus: 200, but Status field: 500,400,200...
func (h HTTPHandler) AsJson(ctx *app.Context, status int, message string, data interface{}) *server.Response {
	return h.App.AsJson(ctx, status, message, data)
}

func (h HTTPHandler) AsJsonInterface(ctx *app.Context, status int, data interface{}) *server.ResponseInterface {
	return h.App.AsJsonInterface(ctx, status, data)
}

// BadRequest For mobile, httpStatus:200, but Status field: http.MobileBadRequest
func (h HTTPHandler) BadRequest(ctx *app.Context, err error) *server.Response {
	return h.App.AsJson(ctx, http.StatusBadRequest, err.Error(), nil)
}

// ForbiddenRequest For mobile, httpStatus:200, but Status field: http.StatusForbidden
func (h HTTPHandler) ForbiddenRequest(ctx *app.Context, err error) *server.Response {
	return h.App.AsJson(ctx, http.StatusForbidden, err.Error(), nil)
}

func (h HTTPHandler) AsError(ctx *app.Context, message string) *server.Response {
	return h.App.AsJson(ctx, http.StatusInternalServerError, message, nil)
}

func (h HTTPHandler) ThrowBadRequestException(ctx *app.Context, message string) *server.Response {
	return h.App.ThrowExceptionJson(ctx, http.StatusBadRequest, 0, "Bad Request", message)
}

func (h HTTPHandler) CreateUser(ctx *gin.Context) {
	body := domain.User{
		Username: ctx.PostForm("username"),
		Email:    ctx.PostForm("email"),
		Password: ctx.PostForm("password"),
	}
	// if err := ctx.ShouldBind(&body); err != nil {
	// 	return h.AsJsonInterface(ctx, http.StatusBadRequest, err)
	// }
	if bodyCheck := body.CheckData(); bodyCheck != "" {
		h.AsDataUnauthorized(ctx, bodyCheck)
		return
	}

	if duplicateCheck, err := h.UserService.Login(ctx, body.Username, body.Email); err != nil {
		h.AsDatabaseError(ctx)
		return
	} else if duplicateCheck.Email == body.Email && duplicateCheck.Username == body.Username {
		h.AsDuplicateEmail(ctx)
		return
	}

	if err := body.HashPassword(body.Password); err != nil {
		h.AsHashError(ctx)
		return
	}

	if err := h.UserService.CreateUser(ctx, &body); err != nil {
		h.AsDatabaseError(ctx)
		return
	}
	verification := domain.Verification{
		UserId:   body.Id,
		Verified: false,
	}
	if err := h.UserService.CreateVerification(ctx, &verification); err != nil {
		h.AsDatabaseError(ctx)
		return
	}

	if err := mail.Verify_mail(&body); err != nil {
		h.AsEmailNotFound(ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Account created, check your email for verification",
		"Id":      body.Id.String(),
	})
}

func (h HTTPHandler) Login(ctx *gin.Context) {
	//Declaring Variables
	var tokenString string
	var verified *bool
	request := domain.UserLogin{
		Username: ctx.PostForm("username"),
		Email:    ctx.PostForm("email"),
		Password: ctx.PostForm("password"),
	}
	resp, err := h.UserService.Login(ctx, request.Username, request.Email)
	if err != nil {
		h.AsDatabaseError(ctx)
		return

	}
	if resp.Id == uuid.Nil {
		h.AsDataNotFound(ctx)
		return
	}
	if err := resp.CheckPassword(request.Password); err != nil {
		h.AsPasswordUnmatched(ctx)
		return
	}
	verified, err = h.UserService.CheckVerified(ctx, resp.Id)
	if err != nil {
		h.AsDatabaseError(ctx)
		return

	}
	if !*verified {
		h.AsNotVerfied(ctx)
		return
	}
	checkJWT, err := h.UserService.CheckJWT(ctx, resp.Id)
	if err != nil {
		h.AsDatabaseError(ctx)
		return
	} else if checkJWT.Jwt == "" {
		tokenString, err = jwthelper.GenerateJWT(*resp, *verified)
		if err != nil {
			h.AsInvalidTokenError(ctx)
			return
		}
		if err := h.UserService.StoreJWT(ctx, tokenString, resp.Id); err != nil {
			h.AsDataNotFound(ctx)
			return
		}
	} else if checkJWT.Jwt != "" {
		tokenString, err = jwthelper.GenerateJWT(*resp, *verified)
		if err != nil {
			h.AsInvalidTokenError(ctx)
			return
		}
		JWTRead, err := jwthelper.TokenRead(checkJWT.Jwt)
		if err != nil {
			h.AsInvalidTokenError(ctx)
			return
		}
		if JWTRead.LastLogin.Day() == time.Now().Day() {
			h.AsJWTExist(ctx)
			return
		} else {
			if err := h.UserService.StoreJWT(ctx, tokenString, resp.Id); err != nil {
				h.AsDataNotFound(ctx)
				return
			}
			ctx.JSON(http.StatusOK, gin.H{
				"token": tokenString,
			})
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

func (h HTTPHandler) UpdateVerification(ctx *gin.Context) {
	idParam := ctx.Param("id")
	Id, err := uuid.Parse(idParam)
	if err != nil {
		h.AsRequiredClientIdError(ctx)
		return
	}
	if err := h.UserService.UpdateVerification(ctx, Id); err != nil {
		h.AsDataNotFound(ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Account verified, please relogin",
	})
}
