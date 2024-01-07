package domain

type EventRepository interface {
	GetAll(userID string) ([]Event, error)
	Create(event *Event) error
	Delete(eventID string) error
	Update(event *Event) (Event, error)
}
