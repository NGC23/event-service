package main

import (
	"database/sql"
	"event-service/internal/application"
	"event-service/internal/infrastructure"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
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

	conn, err := sql.Open("mysql", "neil:gym4life@tcp(jeeves-mysql:3306)/jeeves_api")

	if err != nil {
		panic("we will fix this, we dont panic here")
	}

	defer conn.Close()

	er := infrastructure.NewEventsRepository(conn)
	ec := application.EventController{
		Repository: er,
	}

	router := gin.Default()

	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     false,
		ValidateHeaders: false,
	}))

	router.GET("/events/:userID", ec.GetAll)
	router.GET("/events/details/:ID", ec.GetByID)
	router.POST("/events/create", ec.Create)
	router.PUT("/events/update/:ID", ec.Delete)
	router.DELETE("/events/delete/:ID", ec.Delete)
	router.GET("/health", healthHandler)
	router.GET("/readiness", readinessHandler)

	router.Run(":8081")
	fmt.Println("Server running on port 8081 neil 2")
}
