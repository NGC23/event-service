package infrastructure

import (
	"database/sql"
	"event-service/internal/domain"

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

func (r *eventRepository) Create(e *domain.Event) (*domain.Event, error) {
	_, err := r.conn.Exec("INSERT INTO `events` (`uuid`, `name`, `description`, `created_at`, `start_date`, `end_date`, `user_id`) VALUES(?,?,?,?,?,?,?)", e.ID, e.Name, e.Description, e.CreatedAt, e.StartDate, e.EndDate, e.UserID)

	if err != nil {
		return nil, err
	}

	return e, nil
}

func (r *eventRepository) GetAll(userID string) ([]domain.Event, error) {
	var events []domain.Event

	result, err := r.conn.Query("SELECT * FROM `events` WHERE `user_id`=?", userID)

	if err != nil {
		return nil, err
	}

	defer result.Close()

	for result.Next() {
		var e domain.Event

		err := result.Scan(&e.ID, &e.Name, &e.Description, &e.CreatedAt, &e.StartDate, &e.EndDate, &e.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, e)
	}

	return events, nil
}

func (r *eventRepository) GetByID(ID string) (domain.Event, error) {
	var e domain.Event

	result := r.conn.QueryRow("SELECT * FROM `events` WHERE `uuid`=?", ID)

	err := result.Scan(&e.ID, &e.Name, &e.Description, &e.CreatedAt, &e.StartDate, &e.EndDate, &e.UserID)

	if err != nil {
		return e, err
	}

	return e, nil
}

func (r *eventRepository) Delete(ID string) error {
	_, err := r.conn.Query("DELETE FROM `events` WHERE `uuid`=?", ID)

	if err != nil {
		return err
	}

	return nil
}

func (r *eventRepository) Update(e *domain.Event) error {
	_, err := r.conn.Exec("UPDATE `events` SET name=?, description=?, start_date=?, end_date=? WHERE id=?", e.Name, e.Description, e.StartDate, e.EndDate, e.ID)

	if err != nil {
		return err
	}

	return nil
}
