package domain

type EventRepository interface {
	GetAll(userID string) ([]Event, error)
	GetByID(ID string) (Event, error)
	Create(e *Event) (*Event, error)
	Delete(eventID string) error
	Update(event *Event) error
}
