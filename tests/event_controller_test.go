package tests

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"event-service/internal/application"
	"event-service/internal/domain"
	"event-service/internal/infrastructure"
	"log"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateWillFailWith400WhenBodyIsEmpty(t *testing.T) {
	db, _, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database", err)
	}

	defer db.Close()

	er := infrastructure.NewEventsRepository(db)
	ec := application.EventController{
		Repository: er,
	}

	w := httptest.NewRecorder()
	_, engine := gin.CreateTestContext(w)
	engine.POST("/events/create", ec.Create)
	req, _ := http.NewRequest("POST", "/events/create", nil)
	engine.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 400)
}

func TestCreateOK(t *testing.T) {
	db, mock := NewMock()

	e := domain.Event{
		Name:        "blah",
		UserID:      "random-user-id",
		StartDate:   "2024-01-14 20:08:56",
		EndDate:     "2024-01-15 20:08:56",
		CreatedAt:   "2024-01-14 20:08:56",
		Description: "This is a test",
	}

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `events` (`uuid`, `name`, `description`, `start_date`, `end_date`, `user_id`) VALUES(?,?,?,?,?,?)")).WithArgs(sqlmock.AnyArg(), e.Name, e.Description, e.StartDate, e.EndDate, e.UserID).WillReturnResult(sqlmock.NewResult(0, 0))

	er := infrastructure.NewEventsRepository(db)
	ec := &application.EventController{Repository: er}

	json, _ := json.Marshal(e)

	w := httptest.NewRecorder()
	_, engine := gin.CreateTestContext(w)
	engine.POST("/events/create", ec.Create)
	req, _ := http.NewRequest("POST", "/events/create", bytes.NewBuffer((json)))
	engine.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 201)
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}
