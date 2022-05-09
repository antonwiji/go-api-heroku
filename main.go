package main

import (
	"fmt"
	"os"
	"rest-api-gorm/handler"
	"rest-api-gorm/middlerwer"
	"rest-api-gorm/repository"
	"rest-api-gorm/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	// err := godotenv.Load(".env")

	// if err != nil {
	// 	panic(err)
	// }

	host := os.Getenv("HOST")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbName := os.Getenv("DBNAME")
	dbPort := os.Getenv("DBPORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=Asia/Shanghai", host, user, password, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.Use(middlerwer.CORSMiddleware())

	NewRepo := repository.NewRepository(db)
	NewService := service.NewService(NewRepo)
	NewHandler := handler.NewHandler(NewService)

	r.GET("/user", NewHandler.FindAll)
	r.GET("/user/:id", NewHandler.FindById)
	r.PUT("/user/:id", NewHandler.Update)
	r.DELETE("/user/:id", NewHandler.Delate)
	r.POST("/user", NewHandler.CreateApi)

	r.Run()

}
