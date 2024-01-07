package domain

type Event struct {
	ID          string `json:"id,omitempty" db:"uuid"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	StartDate   string `json:"start_date"  db:"start_date"`
	EndDate     string `json:"end_date" db:"end_date"`
	CreatedAt   string `json:"created_at,omitempty"`
	UserID      string `json:"user_id" db:"user_id"`
}
