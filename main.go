package main

import (
	auth "job-com-backend/authentication"
	config "job-com-backend/config"
	controllers "job-com-backend/controllers"
	routers "job-com-backend/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	var secretKey string = "8a9b4555-664d-4afd-8598-a8d57fe6ab7d"
	jwtAuth, _ := auth.NewJWTTokenMaker(secretKey)
	controllers.Initialize(config.Connect(), jwtAuth)
	router := gin.Default()
	routers.Routers(router)
	router.Run(":8080")
}
