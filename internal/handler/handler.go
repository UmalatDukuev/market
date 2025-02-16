package handler

import (
	"market/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	api := router.Group("/api")
	// api := router.Group("/api", h.userIdentity)
	// {
	// 	api.GET("/info", h.getInfo)
	// }
	auth := api.Group("/auth")
	{
		auth.POST("sign-up", h.signUp)
		auth.POST("sign-in", h.signIn)
	}

	api.GET("/info", h.userIdentity, h.getInfo)
	api.POST("/sendCoin", h.userIdentity, h.sendCoin)

	return router
}
