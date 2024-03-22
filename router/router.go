package router

import (
	"MyGram/controllers"
	"MyGram/middleware"
	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
		userRouter.PUT("/", middleware.Authentication(), controllers.UserUpdate)
		userRouter.DELETE("/", middleware.Authentication(), controllers.UserDelete)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middleware.Authentication())
		photoRouter.POST("/", controllers.AddPhoto)
		photoRouter.GET("/", controllers.GetPhotos)
		photoRouter.PUT("/:id", middleware.PhotoAuthorization(), controllers.UpdatePhoto)
		photoRouter.DELETE("/:id", middleware.PhotoAuthorization(), controllers.DeletePhoto)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middleware.Authentication())
		commentRouter.POST("/", controllers.AddComment)
		commentRouter.GET("/", controllers.GetComments)
		commentRouter.PUT("/:id", middleware.CommentAuthorization(), controllers.UpdateComment)
		commentRouter.DELETE("/:id", middleware.CommentAuthorization(), controllers.DeleteComment)
	}

	socialMediaRouter := r.Group("/socialmedias")
	{
		socialMediaRouter.Use(middleware.Authentication())
		socialMediaRouter.POST("/", controllers.AddSocialMedia)
		socialMediaRouter.GET("/", controllers.GetSocialMedias)
		socialMediaRouter.PUT("/:id", middleware.SocialMediaAuthorization(), controllers.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:id", middleware.SocialMediaAuthorization(), controllers.DeleteSocialMedia)
	}

	testRouter := r.Group("/test")
	{
		testRouter.GET("/", middleware.Authentication(), controllers.GetAllUsers)
		testRouter.GET("/me", middleware.Authentication(), controllers.UserMe)
	}

	return r
}
