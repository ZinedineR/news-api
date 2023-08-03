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
	// h.MoodleRoute("POST", "/user", h.userHandler.CreateUser)
	// h.MoodleRoute("GET", "/user", h.userHandler.GetUserData)
	// h.MoodleRoute("GET", "/verify/:id", h.userHandler.UpdateVerification)
}

func (h *HttpServe) MoodleRoute(method, path string, f handler.HandlerFnInterface) {
	switch method {
	case "GET":
		h.router.GET(path, h.base.MoodleRunAction(f))
	case "POST":
		h.router.POST(path, h.base.MoodleRunAction(f))
	case "PUT":
		h.router.PUT(path, h.base.MoodleRunAction(f))
	case "DELETE":
		h.router.DELETE(path, h.base.MoodleRunAction(f))
	default:
		panic(fmt.Sprintf(":%s method not allow", method))
	}
}
