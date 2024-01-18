package test

import (
	"buyer/src/adapter"
	"buyer/src/datasource"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// before all
func init() {
}

func TestRouter(t *testing.T) {
	database := datasource.NewDatabase("m")
	database.Connect()

	router := adapter.NewRouter()
	server := adapter.NewServer(router)
	server.Routing()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/users", nil)
	server.Router.Engine.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"users\":[]}", w.Body.String())

	jsonData := []byte(`{"user_age": 15, "user_name": "sewon"}`)
	body := bytes.NewBuffer(jsonData)
	req, _ = http.NewRequest("POST", "/api/v1/user", body)
	req.Header.Set("Content-Type", "application/json") // JSON 컨텐츠 타입 설정

	w = httptest.NewRecorder()
	server.Router.Engine.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"age\":15,\"name\":\"sewon\"}", w.Body.String())

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/v1/users", nil)
	server.Router.Engine.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"users":[{"id":1,"age":15,"name":"sewon"}]}`, w.Body.String())
}
