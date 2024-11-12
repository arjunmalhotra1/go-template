package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arjunmalhotra1/go-template/internal/handler"
	"github.com/stretchr/testify/assert"
)

func TestCreateTask(t *testing.T) {
	req := httptest.NewRequest("POST", "/tasks", nil)
	w := httptest.NewRecorder()
	handler := handler.NewTaskHandler(nil)

	handler.CreateTask(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
