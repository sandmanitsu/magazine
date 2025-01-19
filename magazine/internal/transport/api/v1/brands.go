package v1

import (
	service "magazine/internal/services"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func (h *Hanlder) initBrandRoutes(api *gin.RouterGroup) {
	items := api.Group("/brand")
	{
		items.GET("/list", h.brandsList)
		items.POST("/signup", h.signUpBrand)
		items.POST("/signin", h.signInBrand)
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

type BrandSignUpData struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Login    string `json:"login"`
}

func (h *Hanlder) signUpBrand(c *gin.Context) {
	var data BrandSignUpData
	err := c.BindJSON(&data)
	if err != nil {
		c.String(http.StatusBadRequest, "Error: invalid data")

		return
	}

	err = h.services.Brand.SignUp(service.BrandSignUpData{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		Login:    data.Login,
	})
	if err != nil {
		c.String(http.StatusBadRequest, "Failed")

		return
	}

	c.String(http.StatusOK, "Success")
}

type BrandSignInData struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (h *Hanlder) signInBrand(c *gin.Context) {
	var data BrandSignInData
	err := c.BindJSON(&data)
	if err != nil {
		c.String(http.StatusBadRequest, "Error: invalid data")

		return
	}

	tokens, err := h.services.Brand.SignIn(service.BrandSignInData{
		Login:    data.Login,
		Password: data.Password,
	})
	if err != nil {
		c.String(http.StatusBadRequest, "Error: invalid data")

		return
	}

	c.JSON(http.StatusOK, tokens)
}
