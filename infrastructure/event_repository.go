package infrastructure

import (
	"database/sql"
	"event-service/domain"
	"fmt"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

// eventRepository Implements domain.EventRepository
type eventRepository struct {
	conn *sql.DB
}

// NewEventsRepository returns initialized EventRepositoryInterface
func NewEventsRepository(conn *sql.DB) domain.EventRepository {
	return &eventRepository{conn: conn}
}

func (r *eventRepository) Create(context *gin.Context, event *domain.Event) error {
	event.Id = uuid.NewString()

	_, err := r.conn.Exec(fmt.Sprintf("INSERT INTO `events` VALUES('%s', '%s', '%s', '%s', '%s', '%s')", event.Id, event.Name, event.Description, event.StartDate, event.EndDate, event.UserId))

	if err != nil {
		return err
	}

	return nil
}

func (r *eventRepository) GetAll(context *gin.Context) ([]domain.Event, error) {
	var events []domain.Event

	// TODO: make this user id specific context and add the index to user table for user_id
	result, err := r.conn.Query("SELECT * FROM `events`")

	if err != nil {
		return nil, err
	}

	defer result.Close()

	for result.Next() {
		var event domain.Event
		if err := result.Scan(&event.Id, &event.Name, &event.Description, &event.StartDate, &event.EndDate, &event.UserId); err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func (r *eventRepository) Delete(context *gin.Context) error {
	return nil
}
