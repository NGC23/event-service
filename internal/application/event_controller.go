package application

import (
	"event-service/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

// EventController provides use-case
type EventController struct {
	Repository domain.EventRepository
}

func (c EventController) Create(ctx *gin.Context) {
	var event domain.Event

	err := ctx.ShouldBindJSON(&event)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	e, err := c.Repository.Create(&event)

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusCreated, e)
}

func (c EventController) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, nil)
}

func (c EventController) GetAll(ctx *gin.Context) {
	userID := ctx.Param("userID")

	if userID == "" {
		//TODO: see if response is suitable for now just a bad request
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	result, err := c.Repository.GetAll(userID)

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (c EventController) GetByID(ctx *gin.Context) {
	ID := ctx.Param("ID")

	if ID == "" {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	result, err := c.Repository.GetByID(ID)

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (c EventController) Delete(ctx *gin.Context) {
	ID := ctx.Param("ID")

	if ID == "" {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err := c.Repository.Delete(ID)

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
