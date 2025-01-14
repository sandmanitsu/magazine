package v1

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func (h *Hanlder) initItemsRoutes(api *gin.RouterGroup) {
	items := api.Group("/items")
	{
		items.GET("/list", h.itemsList)
	}
}

func (h *Hanlder) itemsList(c *gin.Context) {
	params, err := url.ParseQuery(c.Request.URL.RawQuery)
	if err != nil {
		c.String(http.StatusBadRequest, "Error: parsing query params")

		return
	}

	item, err := h.services.Items.Items(params)
	if err != nil {
		c.String(http.StatusBadRequest, "Error: getting data")

		return
	}

	c.JSON(http.StatusOK, item)
}
