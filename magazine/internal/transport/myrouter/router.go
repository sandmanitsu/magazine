package myrouter

import (
	service "magazine/internal/services"
	v1 "magazine/internal/transport/api/v1"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Hanlder struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Hanlder {
	return &Hanlder{
		services: services,
	}
}

func (h *Hanlder) Init() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	h.initAPI(router)

	return router
}

func (h *Hanlder) initAPI(router *gin.Engine) {
	hanlderV1 := v1.NewHandler(h.services)
	api := router.Group("/api")
	{
		hanlderV1.Init(api)
	}
}
