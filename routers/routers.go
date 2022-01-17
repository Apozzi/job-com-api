package routers

import (
	controllers "job-com-api/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Routers(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowMethods:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	router.POST("login", controllers.Login)
	router.GET("oportunity/paginated", controllers.GetOportunitiesPaginated)
	router.GET("oportunity/:id", controllers.GetOportunity)
	router.POST("user", controllers.PostUser)
	router.POST("email", controllers.SendEmail)
	privateRoute := router.Group("")
	privateRoute.Use(controllers.Authenticate)
	{
		privateRoute.GET("verifyToken", controllers.VerifyToken)
		privateRoute.GET("oportunity", controllers.GetOportunities)
		privateRoute.GET("oportunity/find-by-user-id/:userId", controllers.FindOportunitiesByUser)
		privateRoute.POST("oportunity", controllers.PostOportunity)
		privateRoute.PUT("oportunity", controllers.PutOportunity)
		privateRoute.DELETE("oportunity/:id", controllers.DeleteOportunity)
		privateRoute.GET("user", controllers.GetUsers)
		privateRoute.GET("user/:id", controllers.GetUser)
		privateRoute.PUT("user", controllers.PutUser)
		privateRoute.DELETE("user/:id", controllers.DeleteUser)
	}
}
