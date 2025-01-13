package v1

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func (h *Hanlder) initBrandRoutes(api *gin.RouterGroup) {
	items := api.Group("/brand")
	{
		items.GET("/list", h.brandsList)
	}
}

func (h *Hanlder) brandsList(c *gin.Context) {
	params, err := url.ParseQuery(c.Request.URL.RawQuery)
	if err != nil {
		c.String(http.StatusBadRequest, "Error: parsing query params")
	}

	brands, err := h.services.Brand.Brands(params)
	if err != nil {
		c.String(http.StatusBadRequest, "Error: getting data")
	}

	c.JSON(http.StatusOK, brands)
}
