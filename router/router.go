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
		photoRouter.GET("/:photo_id", controller.GetPhotoByID)
		photoRouter.Use(middleware.Authentication())
		photoRouter.POST("/", controller.CreatePhoto)
		photoRouter.GET("/", controller.GetMyPhoto)
		photoRouter.PUT("/:photo_id", middleware.PhotoAuthorization(), controller.EditPhoto)
		photoRouter.DELETE("/:photo_id", middleware.PhotoAuthorization(), controller.PhotoDelete)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middleware.Authentication())
		commentRouter.POST("/", controller.CreateComment)
		commentRouter.GET("/", controller.GetMyComment)
		commentRouter.PUT("/:comment_id", middleware.CommentAuthorization(), controller.EditComment)
		commentRouter.DELETE("/:comment_id", middleware.CommentAuthorization(), controller.DeleteComment)
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
