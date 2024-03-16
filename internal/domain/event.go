package domain

type Event struct {
	ID          string `json:"id,omitempty" db:"uuid"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	CreatedAt   string `json:"created_at,omitempty"  db:"created_at"`
	StartDate   string `json:"start_date"  db:"start_date"`
	EndDate     string `json:"end_date" db:"end_date"`
	UserID      string `json:"user_id" db:"user_id"`
}
