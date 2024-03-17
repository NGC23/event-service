package tests

import (
	"bytes"
	"encoding/json"
	"event-service/internal/application"
	"event-service/internal/domain"
	"event-service/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateWillFailWith400WhenBodyIsEmpty(t *testing.T) {
	erm := new(mocks.EventRepository)
	ec := &application.EventController{Repository: erm}

	w := httptest.NewRecorder()

	_, engine := gin.CreateTestContext(w)
	engine.POST("/events/create", ec.Create)

	req, _ := http.NewRequest(http.MethodPost, "/events/create", nil)
	engine.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreateOK(t *testing.T) {
	erm := new(mocks.EventRepository)
	ec := &application.EventController{Repository: erm}

	e := domain.Event{
		ID:          "",
		Name:        "blah",
		UserID:      "random-user-id",
		StartDate:   "2024-01-14 20:08:56",
		EndDate:     "2024-01-15 20:08:56",
		CreatedAt:   "2024-01-14 20:08:56",
		Description: "This is a test",
	}

	erm.On("Create", mock.Anything).Return(nil)

	json, _ := json.Marshal(e)

	w := httptest.NewRecorder()

	_, engine := gin.CreateTestContext(w)
	engine.POST("/events/create", ec.Create)

	req, _ := http.NewRequest(http.MethodPost, "/events/create", bytes.NewBuffer((json)))
	engine.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestDeleteOk(t *testing.T) {
	erm := new(mocks.EventRepository)
	ec := &application.EventController{Repository: erm}

	erm.On("Delete", mock.Anything).Return(nil)

	w := httptest.NewRecorder()

	_, engine := gin.CreateTestContext(w)
	engine.DELETE("/events/delete/:ID", ec.Delete)

	req, _ := http.NewRequest(http.MethodDelete, "/events/delete/user-test-id", nil)
	engine.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestDeleteNotOkWhenNoIDProvided(t *testing.T) {
	erm := new(mocks.EventRepository)
	ec := &application.EventController{Repository: erm}

	w := httptest.NewRecorder()

	_, engine := gin.CreateTestContext(w)
	engine.DELETE("/events/delete/:ID", ec.Delete)

	req, _ := http.NewRequest(http.MethodDelete, "/events/delete/", nil)
	engine.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
