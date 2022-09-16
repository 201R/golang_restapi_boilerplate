package v1

import (
	"net/http"

	"github.com/201R/go_api_boilerplate/packages/app"
	"github.com/201R/go_api_boilerplate/services"
	"github.com/gin-gonic/gin"
)

// UserController interface
type UserController interface {
	Get(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
	GetByID(*gin.Context)
}

type userController struct {
	us services.UserService
}

func NewUserController(us services.UserService) UserController {
	return userController{
		us: us,
	}
}

// Create implements UserController
func (ctl userController) Create(c *gin.Context) {
	app.HTTPRes(c, http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented), nil)
}

// @Summary Remove user given id
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/user/{id} [delete]
// @tags User
func (ctl userController) Delete(c *gin.Context) {
	app.HTTPRes(c, http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented), nil)
}

// Get implements UserController
func (ctl userController) Get(c *gin.Context) {
	app.HTTPRes(c, http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented), nil)
}

// @Summary Get user info of given id
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/user/{id} [get]
// @tags User
func (ctl userController) GetByID(c *gin.Context) {
	app.HTTPRes(c, http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented), nil)
}

func (ctl userController) Update(c *gin.Context) {
	app.HTTPRes(c, http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented), nil)
}
