package handler

import (
	"net/http"

	"news-api/internal/base/app"
	BaseDomain "news-api/internal/base/domain"
	"news-api/internal/base/handler"
	"news-api/internal/news/domain"
	NewsService "news-api/internal/news/service"
	"news-api/pkg/responsehelper"
	"news-api/pkg/server"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type HTTPHandler struct {
	App         *handler.BaseHTTPHandler
	NewsService NewsService.Service
}

func NewHTTPHandler(handler *handler.BaseHTTPHandler, NewsService NewsService.Service) *HTTPHandler {
	return &HTTPHandler{
		App:         handler,
		NewsService: NewsService,
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
		"responseMessage": "Another account with same email already created",
	})
}

func (h HTTPHandler) AsDataNotFound(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, gin.H{
		"responseCode":    "404",
		"responseMessage": "Data not Found",
	})
}

func (h HTTPHandler) AsJWTExist(ctx *gin.Context, token string) {
	ctx.JSON(http.StatusUnauthorized, gin.H{
		"responseCode":    "401",
		"responseMessage": "You already login before",
		"token":           token,
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

func (h HTTPHandler) CreateCategories(ctx *app.Context) *server.ResponseInterface {
	body := domain.Categories{
		Title: ctx.PostForm("title"),
	}
	if err := h.NewsService.CreateCategories(ctx, &body); err != nil {
		respStatus := responsehelper.GetStatusResponse(http.StatusBadRequest, "Error in creating categories")
		return h.AsJsonInterface(ctx, http.StatusBadRequest, respStatus)
	}
	respStatus := responsehelper.GetStatusResponse(http.StatusOK, "")

	finalResponse := struct {
		*BaseDomain.Status
		*domain.Categories
	}{respStatus, &body}
	return h.AsJsonInterface(ctx, http.StatusOK, finalResponse)
}

func (h HTTPHandler) GetDetailCategories(ctx *app.Context) *server.ResponseInterface {
	idParam := ctx.Param("id")
	Id, err := uuid.Parse(idParam)
	if err != nil {
		respStatus := responsehelper.GetStatusResponse(http.StatusBadRequest, "Id param not valid")
		return h.AsJsonInterface(ctx, http.StatusBadRequest, respStatus)
	}
	resp, err := h.NewsService.GetDetailCategories(ctx, Id)
	if err != nil {
		respStatus := responsehelper.GetStatusResponse(http.StatusBadRequest, "Data not found/error getting data from database")
		return h.AsJsonInterface(ctx, http.StatusBadRequest, respStatus)
	}
	respStatus := responsehelper.GetStatusResponse(http.StatusOK, "")

	finalResponse := struct {
		*BaseDomain.Status
		*domain.Categories
	}{respStatus, resp}
	return h.AsJsonInterface(ctx, http.StatusOK, finalResponse)
}

func (h HTTPHandler) ListCategories(ctx *app.Context) *server.ResponseInterface {
	resp, err := h.NewsService.GetCategories(ctx)
	if err != nil {
		respStatus := responsehelper.GetStatusResponse(http.StatusBadRequest, "Data not found/error getting data from database")
		return h.AsJsonInterface(ctx, http.StatusBadRequest, respStatus)
	}
	respStatus := responsehelper.GetStatusResponse(http.StatusOK, "")

	finalResponse := struct {
		*BaseDomain.Status
		List []domain.Categories `json:"list"`
	}{respStatus, *resp}
	return h.AsJsonInterface(ctx, http.StatusOK, finalResponse)
}

func (h HTTPHandler) UpdateCategories(ctx *app.Context) *server.ResponseInterface {
	idParam := ctx.Param("id")
	Id, err := uuid.Parse(idParam)
	if err != nil {
		respStatus := responsehelper.GetStatusResponse(http.StatusBadRequest, "Id param not valid")
		return h.AsJsonInterface(ctx, http.StatusBadRequest, respStatus)
	}
	body := domain.Categories{
		Id:    Id,
		Title: ctx.PostForm("title"),
	}
	resp, err := h.NewsService.SearchCategories(ctx, body.Title)
	if err != nil {
		respStatus := responsehelper.GetStatusResponse(http.StatusBadRequest, "Data not found/error getting data from database")
		return h.AsJsonInterface(ctx, http.StatusBadRequest, respStatus)
	}
	if resp.Title == body.Title {
		respStatus := responsehelper.GetStatusResponse(http.StatusUnauthorized, `Category '`+body.Title+`' already in database`)
		return h.AsJsonInterface(ctx, http.StatusUnauthorized, respStatus)
	}
	if err := h.NewsService.UpdateCategories(ctx, &body); err != nil {
		respStatus := responsehelper.GetStatusResponse(http.StatusBadRequest, "Error in creating categories")
		return h.AsJsonInterface(ctx, http.StatusBadRequest, respStatus)
	}
	respStatus := responsehelper.GetStatusResponse(http.StatusOK, "")

	finalResponse := struct {
		*BaseDomain.Status
		*domain.Categories
	}{respStatus, &body}
	return h.AsJsonInterface(ctx, http.StatusOK, finalResponse)
}

func (h HTTPHandler) DeleteCategories(ctx *app.Context) *server.ResponseInterface {
	idParam := ctx.Param("id")
	Id, err := uuid.Parse(idParam)
	if err != nil {
		respStatus := responsehelper.GetStatusResponse(http.StatusBadRequest, "Id param not valid")
		return h.AsJsonInterface(ctx, http.StatusBadRequest, respStatus)
	}

	if err := h.NewsService.DeleteCategories(ctx, Id); err != nil {
		respStatus := responsehelper.GetStatusResponse(http.StatusBadRequest, "Error in deleting categories")
		return h.AsJsonInterface(ctx, http.StatusBadRequest, respStatus)
	}

	respStatus := responsehelper.GetStatusResponse(http.StatusOK, "")

	finalResponse := struct {
		*BaseDomain.Status
		AdditionalInfo string `json:"additional_info"`
	}{respStatus, Id.String() + " has been deleted"}
	return h.AsJsonInterface(ctx, http.StatusOK, finalResponse)
}

func (h HTTPHandler) CreateNews(ctx *app.Context) *server.ResponseInterface {
	categoryForm := ctx.PostForm("categories_id")
	category, err := uuid.Parse(categoryForm)
	if err != nil {
		respStatus := responsehelper.GetStatusResponse(http.StatusBadRequest, "Id param not valid")
		return h.AsJsonInterface(ctx, http.StatusBadRequest, respStatus)
	}
	body := domain.News{
		CategoriesId: category,
		Title:        ctx.PostForm("title"),
		Description:  ctx.PostForm("description"),
		Content:      ctx.PostForm("content"),
	}
	if bodyCheck := body.CheckData(); bodyCheck != "" {
		respStatus := responsehelper.GetStatusResponse(http.StatusUnauthorized, bodyCheck)
		return h.AsJsonInterface(ctx, http.StatusUnauthorized, respStatus)
	}
	if err := h.NewsService.CreateNews(ctx, &body); err != nil {
		respStatus := responsehelper.GetStatusResponse(http.StatusBadRequest, "Error in creating news")
		return h.AsJsonInterface(ctx, http.StatusBadRequest, respStatus)
	}
	respStatus := responsehelper.GetStatusResponse(http.StatusOK, "")

	finalResponse := struct {
		*BaseDomain.Status
		*domain.News
	}{respStatus, &body}
	return h.AsJsonInterface(ctx, http.StatusOK, finalResponse)
}

func (h HTTPHandler) GetDetailNews(ctx *app.Context) *server.ResponseInterface {
	idParam := ctx.Param("id")
	Id, err := uuid.Parse(idParam)
	if err != nil {
		respStatus := responsehelper.GetStatusResponse(http.StatusBadRequest, "Id param not valid")
		return h.AsJsonInterface(ctx, http.StatusBadRequest, respStatus)
	}
	resp, err := h.NewsService.GetDetailNews(ctx, Id)
	if err != nil {
		respStatus := responsehelper.GetStatusResponse(http.StatusBadRequest, "Data not found/error getting data from database")
		return h.AsJsonInterface(ctx, http.StatusBadRequest, respStatus)
	}
	respStatus := responsehelper.GetStatusResponse(http.StatusOK, "")

	finalResponse := struct {
		*BaseDomain.Status
		*domain.News
	}{respStatus, resp}
	return h.AsJsonInterface(ctx, http.StatusOK, finalResponse)
}

func (h HTTPHandler) ListNews(ctx *app.Context) *server.ResponseInterface {
	resp, err := h.NewsService.GetNews(ctx)
	if err != nil {
		respStatus := responsehelper.GetStatusResponse(http.StatusBadRequest, "Data not found/error getting data from database")
		return h.AsJsonInterface(ctx, http.StatusBadRequest, respStatus)
	}
	respStatus := responsehelper.GetStatusResponse(http.StatusOK, "")

	finalResponse := struct {
		*BaseDomain.Status
		List []domain.News `json:"list"`
	}{respStatus, *resp}
	return h.AsJsonInterface(ctx, http.StatusOK, finalResponse)
}

func (h HTTPHandler) UpdateNews(ctx *app.Context) *server.ResponseInterface {
	idParam := ctx.Param("id")
	Id, err := uuid.Parse(idParam)
	if err != nil {
		respStatus := responsehelper.GetStatusResponse(http.StatusBadRequest, "Id param not valid")
		return h.AsJsonInterface(ctx, http.StatusBadRequest, respStatus)
	}
	categoryForm := ctx.PostForm("categories_id")
	category, err := uuid.Parse(categoryForm)
	if err != nil {
		respStatus := responsehelper.GetStatusResponse(http.StatusBadRequest, "Id param not valid")
		return h.AsJsonInterface(ctx, http.StatusBadRequest, respStatus)
	}
	body := domain.News{
		Id:           Id,
		CategoriesId: category,
		Title:        ctx.PostForm("title"),
		Description:  ctx.PostForm("description"),
		Content:      ctx.PostForm("content"),
	}
	if bodyCheck := body.CheckData(); bodyCheck != "" {
		respStatus := responsehelper.GetStatusResponse(http.StatusUnauthorized, bodyCheck)
		return h.AsJsonInterface(ctx, http.StatusUnauthorized, respStatus)
	}
	if err := h.NewsService.UpdateNews(ctx, &body); err != nil {
		respStatus := responsehelper.GetStatusResponse(http.StatusBadRequest, "Error in updating news")
		return h.AsJsonInterface(ctx, http.StatusBadRequest, respStatus)
	}
	respStatus := responsehelper.GetStatusResponse(http.StatusOK, "")

	finalResponse := struct {
		*BaseDomain.Status
		*domain.News
	}{respStatus, &body}
	return h.AsJsonInterface(ctx, http.StatusOK, finalResponse)
}

func (h HTTPHandler) DeleteNews(ctx *app.Context) *server.ResponseInterface {
	idParam := ctx.Param("id")
	Id, err := uuid.Parse(idParam)
	if err != nil {
		respStatus := responsehelper.GetStatusResponse(http.StatusBadRequest, "Id param not valid")
		return h.AsJsonInterface(ctx, http.StatusBadRequest, respStatus)
	}

	if err := h.NewsService.DeleteNews(ctx, Id); err != nil {
		respStatus := responsehelper.GetStatusResponse(http.StatusBadRequest, "Error in deleting news")
		return h.AsJsonInterface(ctx, http.StatusBadRequest, respStatus)
	}

	respStatus := responsehelper.GetStatusResponse(http.StatusOK, "")

	finalResponse := struct {
		*BaseDomain.Status
		AdditionalInfo string `json:"additional_info"`
	}{respStatus, Id.String() + " has been deleted"}
	return h.AsJsonInterface(ctx, http.StatusOK, finalResponse)
}

func (h HTTPHandler) CreateCustom(ctx *app.Context) *server.ResponseInterface {
	body := domain.Custom{
		CustomUrl:   ctx.PostForm("custom_url"),
		Title:       ctx.PostForm("title"),
		Description: ctx.PostForm("description"),
		Content:     ctx.PostForm("content"),
	}
	if bodyCheck := body.CheckData(); bodyCheck != "" {
		respStatus := responsehelper.GetStatusResponse(http.StatusUnauthorized, bodyCheck)
		return h.AsJsonInterface(ctx, http.StatusUnauthorized, respStatus)
	}
	if err := h.NewsService.CreateCustom(ctx, &body); err != nil {
		respStatus := responsehelper.GetStatusResponse(http.StatusBadRequest, "Error in creating custom page")
		return h.AsJsonInterface(ctx, http.StatusBadRequest, respStatus)
	}
	respStatus := responsehelper.GetStatusResponse(http.StatusOK, "")

	finalResponse := struct {
		*BaseDomain.Status
		*domain.Custom
	}{respStatus, &body}
	return h.AsJsonInterface(ctx, http.StatusOK, finalResponse)
}

func (h HTTPHandler) GetDetailCustom(ctx *app.Context) *server.ResponseInterface {
	idParam := ctx.Param("id")
	Id, err := uuid.Parse(idParam)
	if err != nil {
		respStatus := responsehelper.GetStatusResponse(http.StatusBadRequest, "Id param not valid")
		return h.AsJsonInterface(ctx, http.StatusBadRequest, respStatus)
	}
	resp, err := h.NewsService.GetDetailCustom(ctx, Id)
	if err != nil {
		respStatus := responsehelper.GetStatusResponse(http.StatusBadRequest, "Data not found/error getting data from database")
		return h.AsJsonInterface(ctx, http.StatusBadRequest, respStatus)
	}
	respStatus := responsehelper.GetStatusResponse(http.StatusOK, "")

	finalResponse := struct {
		*BaseDomain.Status
		*domain.Custom
	}{respStatus, resp}
	return h.AsJsonInterface(ctx, http.StatusOK, finalResponse)
}

func (h HTTPHandler) ListCustom(ctx *app.Context) *server.ResponseInterface {
	resp, err := h.NewsService.GetCustom(ctx)
	if err != nil {
		respStatus := responsehelper.GetStatusResponse(http.StatusBadRequest, "Data not found/error getting data from database")
		return h.AsJsonInterface(ctx, http.StatusBadRequest, respStatus)
	}
	respStatus := responsehelper.GetStatusResponse(http.StatusOK, "")

	finalResponse := struct {
		*BaseDomain.Status
		List []domain.Custom `json:"list"`
	}{respStatus, *resp}
	return h.AsJsonInterface(ctx, http.StatusOK, finalResponse)
}

func (h HTTPHandler) UpdateCustom(ctx *app.Context) *server.ResponseInterface {
	idParam := ctx.Param("id")
	Id, err := uuid.Parse(idParam)
	if err != nil {
		respStatus := responsehelper.GetStatusResponse(http.StatusBadRequest, "Id param not valid")
		return h.AsJsonInterface(ctx, http.StatusBadRequest, respStatus)
	}
	body := domain.Custom{
		Id:          Id,
		CustomUrl:   ctx.PostForm("custom_url"),
		Title:       ctx.PostForm("title"),
		Description: ctx.PostForm("description"),
		Content:     ctx.PostForm("content"),
	}
	if bodyCheck := body.CheckData(); bodyCheck != "" {
		respStatus := responsehelper.GetStatusResponse(http.StatusUnauthorized, bodyCheck)
		return h.AsJsonInterface(ctx, http.StatusUnauthorized, respStatus)
	}
	if err := h.NewsService.UpdateCustom(ctx, &body); err != nil {
		respStatus := responsehelper.GetStatusResponse(http.StatusBadRequest, "Error in updating custom page")
		return h.AsJsonInterface(ctx, http.StatusBadRequest, respStatus)
	}
	respStatus := responsehelper.GetStatusResponse(http.StatusOK, "")

	finalResponse := struct {
		*BaseDomain.Status
		*domain.Custom
	}{respStatus, &body}
	return h.AsJsonInterface(ctx, http.StatusOK, finalResponse)
}

func (h HTTPHandler) DeleteCustom(ctx *app.Context) *server.ResponseInterface {
	idParam := ctx.Param("id")
	Id, err := uuid.Parse(idParam)
	if err != nil {
		respStatus := responsehelper.GetStatusResponse(http.StatusBadRequest, "Id param not valid")
		return h.AsJsonInterface(ctx, http.StatusBadRequest, respStatus)
	}

	if err := h.NewsService.DeleteCustom(ctx, Id); err != nil {
		respStatus := responsehelper.GetStatusResponse(http.StatusBadRequest, "Error in deleting custom page")
		return h.AsJsonInterface(ctx, http.StatusBadRequest, respStatus)
	}

	respStatus := responsehelper.GetStatusResponse(http.StatusOK, "")

	finalResponse := struct {
		*BaseDomain.Status
		AdditionalInfo string `json:"additional_info"`
	}{respStatus, Id.String() + " has been deleted"}
	return h.AsJsonInterface(ctx, http.StatusOK, finalResponse)
}
