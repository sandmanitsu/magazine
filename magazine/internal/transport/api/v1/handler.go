package v1

import (
	service "magazine/internal/services"

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

func (h *Hanlder) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initItemsRoutes(v1)
		h.initBrandRoutes(v1)
		h.initSignUpRoutes(v1)
	}
}
