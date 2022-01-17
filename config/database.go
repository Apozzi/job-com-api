package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

// Connecting to db
func Connect() *gorm.DB {
	dsn := "host=" + goDotEnvVariable("POSTGRE_HOST") +
		" user=" + goDotEnvVariable("POSTGRE_USER") +
		" password=" + goDotEnvVariable("POSTGRE_PASSWORD") +
		" dbname=" + goDotEnvVariable("POSTGRE_DBNAME") +
		" port=" + goDotEnvVariable("POSTGRE_PORT") + " sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}
	log.Printf("Connected to db")
	return db
}
