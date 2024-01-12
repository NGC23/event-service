package infrastructure

import (
	"database/sql"
	"event-service/internal/domain"
	"fmt"

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

func (r *eventRepository) Create(event *domain.Event) error {
	event.ID = uuid.NewString()

	_, err := r.conn.Exec(fmt.Sprintf("INSERT INTO `events` VALUES('%s', '%s', '%s', '%s', '%s', '%s')", event.ID, event.Name, event.Description, event.StartDate, event.EndDate, event.UserID))

	if err != nil {
		return err
	}

	return nil
}

func (r *eventRepository) GetAll(userID string) ([]domain.Event, error) {
	var events []domain.Event

	result, err := r.conn.Query(fmt.Sprintf("SELECT * FROM `events` WHERE `user_id`='%s'", userID))

	if err != nil {
		return nil, err
	}

	defer result.Close()

	for result.Next() {
		var event domain.Event

		err := result.Scan(&event.ID, &event.Name, &event.Description, &event.StartDate, &event.EndDate, &event.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func (r *eventRepository) Delete(eventID string) error {
	return nil
}

func (r *eventRepository) Update(event *domain.Event) (domain.Event, error) {
	return domain.Event{}, nil
}
