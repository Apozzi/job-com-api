package config

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func goDotEnvVariable(key string) string {
	return os.Getenv(key)
}

// Connecting to db
func Connect() *gorm.DB {
	dsn := goDotEnvVariable("DATABASE_URL")
	if dsn == "" {
		dsn = "host=" + goDotEnvVariable("POSTGRE_HOST") +
			" user=" + goDotEnvVariable("POSTGRE_USER") +
			" password=" + goDotEnvVariable("POSTGRE_PASSWORD") +
			" dbname=" + goDotEnvVariable("POSTGRE_DBNAME") +
			" port=" + goDotEnvVariable("POSTGRE_PORT") + " sslmode=disable"
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}
	log.Printf("Connected to db")
	return db
}
