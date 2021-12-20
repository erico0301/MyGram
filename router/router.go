package router

import (
	"MyGram/controller"
	"MyGram/middleware"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/login", controller.UserLogin)
		userRouter.POST("/register", controller.UserRegister)

		userRouter.Use(middleware.Authentication())
		userRouter.PUT("/:user_id", middleware.UserAuthorization(), controller.UserEdit)
		userRouter.DELETE("/:user_id", middleware.UserAuthorization(), controller.UserDelete)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middleware.Authentication())
		photoRouter.POST("/", controller.CreatePhoto)
		photoRouter.GET("/")
		photoRouter.GET("/:photo_id")
		photoRouter.PUT("/:photo_id", middleware.PhotoAuthorization(), controller.UpdatePhoto)
		photoRouter.DELETE("/:photo_id", controller.PhotoDelete)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.POST("/")
		commentRouter.GET("/")
		commentRouter.PUT("/:comment_id")
		commentRouter.DELETE("/:comment_id")
	}

	socialmediaRouter := r.Group("/socialmedias")
	{
		socialmediaRouter.POST("/")
		socialmediaRouter.GET("/")
		socialmediaRouter.PUT("/:socialmedia_id")
		socialmediaRouter.DELETE("/:socialmedia_id")
	}

	return r
}
