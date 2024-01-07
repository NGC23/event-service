package application

import (
	"event-service/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

// EventController provides use-case
type EventController struct {
	Repository domain.EventRepository
}

func (c EventController) Create(ctx *gin.Context) {
	var event *domain.Event

	println("body of request is going to be parsed")

	err := ctx.ShouldBindJSON(&event)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	err = c.Repository.Create(ctx, event)

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusCreated, nil)
}

func (c EventController) GetAll(ctx *gin.Context) {
	result, err := c.Repository.GetAll(ctx)

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusCreated, result)
}
