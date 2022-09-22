package v1

import (
	"net/http"

	"github.com/201R/go_api_boilerplate/packages/app"
	"github.com/201R/go_api_boilerplate/services"
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(c *gin.Context)
	Logout(c *gin.Context)
	Refresh(c *gin.Context)
	me(c *gin.Context)
}

type authController struct {
	as services.AuthService
}

func NewAuthController(as services.AuthService) AuthController {
	return authController{
		as: as,
	}
}

// Login implements AuthController
func (authController) Login(c *gin.Context) {
	app.HTTPRes(c, http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented), nil)
}

// Logout implements AuthController
func (authController) Logout(c *gin.Context) {
	app.HTTPRes(c, http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented), nil)
}

// Refresh implements AuthController
func (authController) Refresh(c *gin.Context) {
	app.HTTPRes(c, http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented), nil)
}

// me implements AuthController
func (authController) me(c *gin.Context) {
	app.HTTPRes(c, http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented), nil)
}
