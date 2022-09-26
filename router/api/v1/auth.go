package v1

import (
	"net/http"

	"github.com/201R/go_api_boilerplate/dtos"
	"github.com/201R/go_api_boilerplate/packages/app"
	"github.com/201R/go_api_boilerplate/services"
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(c *gin.Context)
	Logout(c *gin.Context)
	Refresh(c *gin.Context)
	Me(c *gin.Context)
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
// Login godoc
// @Summary Login user
// @Produce  json
// @Param credentials body dtos.AuthInput true "User credentials"
// @Success 200 {object} app.Response{data=dtos.LoginReponse}
// @Failure 401 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/auth/login [post]
// @tags Auth
func (ctl authController) Login(c *gin.Context) {
	ctx := c.Request.Context()
	var authInput dtos.AuthInput

	if err := c.ShouldBindJSON(&authInput); err != nil {
		app.HTTPRes(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	loginResponse, err := ctl.as.Login(ctx, &authInput)
	if err != nil {
		app.HTTPRes(c, http.StatusUnauthorized, err.Error(), nil)
		return
	}

	app.HTTPRes(c, http.StatusCreated, http.StatusText(http.StatusAccepted), loginResponse)
}

// Logout implements AuthController
// Logout godoc
// @Summary Log out an user
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/auth/logout [post]
// @tags Auth
func (authController) Logout(c *gin.Context) {
	app.HTTPRes(c, http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented), nil)
}

// Refresh implements AuthController
// @Summary Refresh auth token
// @Produce  json
// @Success 200 {object} app.Response{data=[]dtos.LoginReponse}
// @Failure 500 {object} app.Response
// @Router /api/v1/auth/refresh [get]
// @tags Auth
// @Security ApiKeyAuth
func (authController) Refresh(c *gin.Context) {
	app.HTTPRes(c, http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented), nil)
}

// me implements AuthController
// @Summary Get authenticated user
// @Produce  json
// @Success 200 {object} app.Response{data=[]ent.User}
// @Failure 500 {object} app.Response
// @Router /api/v1/auth/me [get]
// @tags Auth
// @Security ApiKeyAuth
func (authController) Me(c *gin.Context) {
	app.HTTPRes(c, http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented), nil)
}
