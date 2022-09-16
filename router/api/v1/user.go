package v1

import (
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
func (ctl userController) Create(*gin.Context) {
	panic("unimplemented")
}

// Delete implements UserController
func (ctl userController) Delete(*gin.Context) {
	panic("unimplemented")
}

// Get implements UserController
func (ctl userController) Get(*gin.Context) {
	panic("unimplemented")
}

// GetByID implements UserController
func (ctl userController) GetByID(*gin.Context) {
	panic("unimplemented")
}

// Update implements UserController
func (ctl userController) Update(*gin.Context) {
	panic("unimplemented")
}
