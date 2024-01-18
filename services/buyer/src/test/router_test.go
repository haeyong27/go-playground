package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// before all
func init() {
}

func TestRouter(t *testing.T) {

	// database := datasource.NewDatabase("m")
	// database.Connect()

	// router := adapter.NewRouter()
	// server := adapter.NewServer(router)
	// server.Run()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/userss", nil)
	// server.Router.Engine.ServeHTTP(w, req)
	println(req)

	assert.Equal(t, http.StatusOK, w.Code)
	// assert.Equal(t, 200, w.Code)
}
