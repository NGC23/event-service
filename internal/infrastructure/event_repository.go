package infrastructure

import (
	"database/sql"
	"event-service/internal/domain"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// eventRepository Implements domain.EventRepository
type eventRepository struct {
	conn *sql.DB
}

// NewEventsRepository returns initialized EventRepositoryInterface
func NewEventsRepository(conn *sql.DB) domain.EventRepository {
	return &eventRepository{conn: conn}
}

func (r *eventRepository) Create(ctx *gin.Context, e *domain.Event) (*domain.Event, error) {
	// _, err := r.conn.Exec(fmt.Sprintf("INSERT INTO `events` (`uuid`, `name`, `description`, `start_date`, `end_date`, `user_id`) VALUES('%s', '%s', '%s', '%s', '%s', '%s')", e.ID, e.Name, e.Description, e.StartDate, e.EndDate, e.UserID))
	_, err := r.conn.ExecContext(ctx, "INSERT INTO `events` (`uuid`, `name`, `description`, `start_date`, `end_date`, `user_id`) VALUES(?,?,?,?,?,?)", e.ID, e.Name, e.Description, e.StartDate, e.EndDate, e.UserID)

	if err != nil {
		return nil, err
	}

	return e, nil
}

func (r *eventRepository) GetAll(ctx *gin.Context, userID string) ([]domain.Event, error) {
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

func (r *eventRepository) GetByID(ctx *gin.Context, ID string) (domain.Event, error) {
	var e domain.Event

	result, err := r.conn.Query(fmt.Sprintf("SELECT * FROM `events` WHERE `id`='%s'", ID))

	if err != nil {
		return e, err
	}

	defer result.Close()

	err = result.Scan(&e.ID, &e.Name, &e.Description, &e.StartDate, &e.EndDate, &e.UserID)

	if err != nil {
		return e, err
	}

	return e, nil
}

func (r *eventRepository) Delete(ctx *gin.Context, ID string) error {
	_, err := r.conn.Query(fmt.Sprintf("DELETE FROM `events` WHERE `id`='%s'", ID))

	if err != nil {
		return err
	}

	return nil
}

func (r *eventRepository) Update(ctx *gin.Context, event *domain.Event) error {
	return nil
}
