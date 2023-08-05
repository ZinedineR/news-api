package api

import (
	"fmt"

	"news-api/internal/base/handler"
)

func (h *HttpServe) setupRouter() {
	//authentication API
	v1 := h.router.Group("/authentication")
	v1.POST("/login", h.userHandler.Login)
	v1.POST("/user", h.userHandler.CreateUser)
	v1.GET("/verify/:id", h.userHandler.UpdateVerification)

	//Categories
	h.UserRoute("POST", "/categories", h.newsHandler.CreateCategories)
	h.UserRoute("GET", "/categories", h.newsHandler.ListCategories)
	h.UserRoute("GET", "/categories/:id", h.newsHandler.GetDetailCategories)
	h.UserRoute("PUT", "/categories/:id", h.newsHandler.UpdateCategories)
	h.UserRoute("DELETE", "/categories/:id", h.newsHandler.DeleteCategories)

	//News
	h.UserRoute("POST", "/news", h.newsHandler.CreateNews)
	h.GuestRoute("GET", "/news", h.newsHandler.ListNews)
	h.GuestRoute("GET", "/news/:id", h.newsHandler.GetDetailNews)
	h.UserRoute("PUT", "/news/:id", h.newsHandler.UpdateNews)
	h.UserRoute("DELETE", "/news/:id", h.newsHandler.DeleteNews)

	//Custom
	h.UserRoute("POST", "/custom", h.newsHandler.CreateCustom)
	h.GuestRoute("GET", "/custom", h.newsHandler.ListCustom)
	h.GuestRoute("GET", "/custom/:id", h.newsHandler.GetDetailCustom)
	h.UserRoute("PUT", "/custom/:id", h.newsHandler.UpdateCustom)
	h.UserRoute("DELETE", "/custom/:id", h.newsHandler.DeleteCustom)

	//Comment
	h.GuestRoute("POST", "/comment/:id", h.newsHandler.CreateComment)

}

func (h *HttpServe) UserRoute(method, path string, f handler.HandlerFnInterface) {
	switch method {
	case "GET":
		h.router.GET(path, h.base.UserRunAction(f))
	case "POST":
		h.router.POST(path, h.base.UserRunAction(f))
	case "PUT":
		h.router.PUT(path, h.base.UserRunAction(f))
	case "DELETE":
		h.router.DELETE(path, h.base.UserRunAction(f))
	default:
		panic(fmt.Sprintf(":%s method not allow", method))
	}
}

func (h *HttpServe) GuestRoute(method, path string, f handler.HandlerFnInterface) {
	switch method {
	case "GET":
		h.router.GET(path, h.base.GuestRunAction(f))
	case "POST":
		h.router.POST(path, h.base.GuestRunAction(f))
	default:
		panic(fmt.Sprintf(":%s method not allow", method))
	}
}
