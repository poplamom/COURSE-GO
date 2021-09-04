package routes

import (
	"course-go/config"
	"course-go/controllers"
	"course-go/middleware"

	"github.com/gin-gonic/gin"
)

func Serve(r *gin.Engine){
	db := config.GetDB()
	v1 := r.Group("/api/v1")
	authenticate := middleware.Authenticate().MiddlewareFunc()
	authorize := middleware.Authorize()
	
	authGroup := v1.Group("auth")
	authController := controllers.Auth{DB: db}
	{
		authGroup.POST("/sign-up", authController.Signup)
		authGroup.POST("/sign-in", middleware.Authenticate().LoginHandler)
		authGroup.GET("/profile", authenticate, authController.GetProfile)
		authGroup.PATCH("/profile", authenticate, authController.UpdateProfile)
	}

	usersController := controllers.Users{DB: db}
	usersGroup := v1.Group("users")
	usersGroup.Use(authenticate, authorize)
	{
		usersGroup.GET("", usersController.FindAll)
		usersGroup.POST("", usersController.Create)
		usersGroup.GET("/:id", usersController.FindOne)
		usersGroup.PATCH("/:id", usersController.Update)
		usersGroup.DELETE("/:id", usersController.Delete)
		usersGroup.PATCH("/:id/promote", usersController.Promote)
		usersGroup.PATCH("/:id/demote", usersController.Demote)
	}
	
	articleController := controllers.Articles{DB: db}
	articlesGroup := v1.Group("/articles")
	articlesGroup.GET("", articleController.FindAll)
	articlesGroup.GET("/:id", articleController.FindOne)
	articlesGroup.Use(authenticate, authorize)
	{
		articlesGroup.PATCH("/:id", articleController.Update)
		articlesGroup.DELETE("/:id", articleController.Delete)
		articlesGroup.POST("" ,authenticate,articleController.Create)
	}

	CategoryController := controllers.Categories{DB: db}
	categoriesGroup := v1.Group("/categories")
	categoriesGroup.GET("", CategoryController.FindAll)
	categoriesGroup.GET("/:id", CategoryController.FindOne)
	categoriesGroup.Use(authenticate, authorize)
	{
		categoriesGroup.PATCH("/:id", CategoryController.Update)
		categoriesGroup.DELETE("/:id", CategoryController.Delete)
		categoriesGroup.POST("" ,CategoryController.Create)
	}

	dockersGroup := v1.Group("/containers")
	dockerController := controllers.Dockers{}
	{
		dockersGroup.GET("", dockerController.ListAll)
		dockersGroup.GET("/stop/:id", dockerController.StopContainer)
		dockersGroup.GET("/start/:id", dockerController.StartContainer)
	}
	


}