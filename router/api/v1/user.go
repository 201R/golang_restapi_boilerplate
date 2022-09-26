package v1

import (
	"net/http"

	"github.com/201R/go_api_boilerplate/dtos"
	"github.com/201R/go_api_boilerplate/packages/app"
	"github.com/201R/go_api_boilerplate/services"
	"github.com/gin-gonic/gin"
)

// UserController interface
type UserController interface {
	Get(*gin.Context)
	GetByID(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}

type userController struct {
	us services.UserService
}

func NewUserController(us services.UserService) UserController {
	return userController{
		us: us,
	}
}

// Get implements UserController
// @Summary Get all users
// @Produce  json
// @Success 200 {object} app.Response{data=[]ent.User}
// @Failure 500 {object} app.Response
// @Router /api/v1/user [get]
// @tags User
// @Security ApiKeyAuth
func (ctl userController) Get(c *gin.Context) {
	ctx := c.Request.Context()
	users, err := ctl.us.Get(ctx)
	if err != nil {
		app.HTTPRes(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	app.HTTPRes(c, http.StatusOK, http.StatusText(http.StatusOK), users)
}

// @Summary Get user info of given id
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/user/{id} [get]
// @tags User
// @Security ApiKeyAuth
func (ctl userController) GetByID(c *gin.Context) {
	app.HTTPRes(c, http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented), nil)
}

// @Summary Register user
// @Produce  json
// @Param user body dtos.UserInput true "Add user"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/user [post]
// @tags User
// @Security ApiKeyAuth
func (ctl userController) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var userInput dtos.UserInput

	if err := c.ShouldBindJSON(&userInput); err != nil {
		app.HTTPRes(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	user, err := ctl.us.Create(ctx, userInput)
	if err != nil {
		app.HTTPRes(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	app.HTTPRes(c, http.StatusCreated, http.StatusText(http.StatusCreated), user)
}

func (ctl userController) Update(c *gin.Context) {
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
// @Security ApiKeyAuth
func (ctl userController) Delete(c *gin.Context) {
	app.HTTPRes(c, http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented), nil)
}
