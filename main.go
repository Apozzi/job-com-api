package main

import (
	auth "job-com-api/authentication"
	config "job-com-api/config"
	controllers "job-com-api/controllers"
	routers "job-com-api/routers"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	var secretKey string = "8a9b4555-664d-4afd-8598-a8d57fe6ab7d"
	jwtAuth, _ := auth.NewJWTTokenMaker(secretKey)
	controllers.Initialize(config.Connect(), jwtAuth)
	router := gin.Default()
	routers.Routers(router)
	port := os.Getenv("PORT")
	router.Run(":" + port)
}
