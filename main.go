package main

import (
	"database/sql"
	"event-service/application"
	"event-service/infrastructure"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	//Testing the client
)

func main() {
	serveApplication()
}

func healthHandler(context *gin.Context) {
	context.JSON(http.StatusOK, nil)
}

func readinessHandler(context *gin.Context) {
	context.JSON(http.StatusOK, nil)
}

func serveApplication() {
	router := gin.Default()

	conn, err := sql.Open("mysql", "neil:gym4life@tcp(jeeves-mysql:3306)/jeeves_api")

	if err != nil {
		//how will we kill it >>>> mmmmmmmmm
		panic("we will panic now..i know we must avoid these")
	}

	defer conn.Close()

	er := infrastructure.NewEventsRepository(conn)
	ec := &application.EventController{
		Repository: er,
	}

	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     false,
		ValidateHeaders: false,
	}))

	router.POST("/events/create", ec.Create)
	router.GET("/events/:userId", ec.GetAll)
	router.GET("/health", healthHandler)
	router.GET("/readiness", readinessHandler)

	router.Run(":8081")
	fmt.Println("Server running on port 8081 neil 2")
}
