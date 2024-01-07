package domain

import (
	"github.com/gin-gonic/gin"
)

type EventRepository interface {
	GetAll(context *gin.Context) ([]Event, error)
	Create(context *gin.Context, event *Event) error
	Delete(context *gin.Context) error
}
