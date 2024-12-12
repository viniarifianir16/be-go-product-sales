package config

import (
	"be-go-product-sales/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	dbProvider := GetEnvOrDefault("DB_PROVIDER", "mysql")

	var db *gorm.DB

	if dbProvider == "postgres" {
		username := os.Getenv("DB_USERNAME")
		password := os.Getenv("DB_PASSWORD")
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		database := os.Getenv("DB_NAME")

		// production DSN
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require",
			host, username, password, database, port)
		dbGorm, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err.Error())
		}

		db = dbGorm

	} else {
		username := GetEnvOrDefault("DB_USERNAME", "root")
		password := GetEnvOrDefault("DB_PASSWORD", "root")
		host := GetEnvOrDefault("DB_HOST", "127.0.0.1")
		port := GetEnvOrDefault("DB_PORT", "3306")
		database := GetEnvOrDefault("DB_NAME", "db_product_sales")

		// production DSN
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			username, password, host, port, database)
		dbGorm, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		log.Printf("Connecting to database with: host=%s port=%s user=%s dbname=%s", host, port, username, database)

		if err != nil {
			panic(err.Error())
		}

		db = dbGorm
	}

	db.AutoMigrate(
		&models.Product{},
	)

	return db
}

func GetEnvOrDefault(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
