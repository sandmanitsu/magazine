package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Hanlder) initSignUpRoutes(api *gin.RouterGroup) {
	items := api.Group("/signup")
	{
		items.GET("/brand", h.signUpBrand)
	}
}

func (h *Hanlder) signUpBrand(c *gin.Context) {
	c.String(http.StatusOK, "Hello!")
}
