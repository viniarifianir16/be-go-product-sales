package api

import (
	"be-go-product-sales/config"
	"be-go-product-sales/docs"
	"be-go-product-sales/routes"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	App *gin.Engine
)

func init() {
	App = gin.New()

	environment := GetEnvOrDefault("ENVIRONMENT", "development")

	if environment == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Println("Warning: Error loading .env file, proceeding without it")
		}
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	docs.SwaggerInfo.Title = "Product Sales REST API"
	docs.SwaggerInfo.Description = "This is REST API Product Sales."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = GetEnvOrDefault("HOST", "localhost:8080")
	if environment == "development" {
		docs.SwaggerInfo.Schemes = []string{"http", "https"}
	} else {
		docs.SwaggerInfo.Schemes = []string{"https"}
	}

	db := config.ConnectDatabase()

	routes.SetupRouter(db, App)
}

func GetEnvOrDefault(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func Handler(w http.ResponseWriter, r *http.Request) {
	App.ServeHTTP(w, r)
}
