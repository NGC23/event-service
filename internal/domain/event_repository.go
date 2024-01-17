package domain

import "github.com/gin-gonic/gin"

type EventRepository interface {
	GetAll(ctx *gin.Context, userID string) ([]Event, error)
	GetByID(ctx *gin.Context, ID string) (Event, error)
	Create(ctx *gin.Context, e *Event) (*Event, error)
	Delete(ctx *gin.Context, eventID string) error
	Update(ctx *gin.Context, event *Event) error
}
