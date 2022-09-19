package router

import (
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	_ "github.com/201R/go_api_boilerplate/docs"
	"github.com/201R/go_api_boilerplate/ent"
	"github.com/201R/go_api_boilerplate/repository"
	v1 "github.com/201R/go_api_boilerplate/router/api/v1"
	"github.com/201R/go_api_boilerplate/services"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.New()
)

func InitRouter(entClient ent.Client) *gin.Engine {

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	/*
		====== Setup middlewares ========
	*/
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	/*
		====== Setup repositories =======
	*/
	userRepo := repository.NewUserRepository(entClient.User)

	/*
		====== Setup Service =======
	*/
	userService := services.NewUserService(userRepo)

	/*
		====== Setup Service =======
	*/
	userController := v1.NewUserController(userService)

	/*
		====== Setup routes =============
	*/
	g := router.Group("/api")
	gV1 := g.Group("/v1")
	{
		gUser := gV1.Group("user")
		{
			gUser.GET("", userController.Get)
			gUser.GET(":id", userController.GetByID)
			gUser.POST("", userController.Create)
			gUser.PUT(":id", userController.Update)
			gUser.DELETE(":id", userController.Delete)
		}
	}

	return router
}
