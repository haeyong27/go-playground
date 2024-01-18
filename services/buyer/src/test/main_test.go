package test

import (
	"buyer/src/adapter"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	router := setupRouter()

	t.Run("TestRootEndpoint", func(t *testing.T) {
		w := performRequest(router, "GET", "/")
		assert.Equal(t, http.StatusOK, w.Code)

		expected := `{"message":"Hello, World!"}`
		assert.Equal(t, expected, w.Body.String())
	})

	t.Run("TestUsersEndpoint", func(t *testing.T) {
		w := performRequest(router, "GET", "/users")
		assert.Equal(t, http.StatusOK, w.Code)

		// TODO: Add assertions for the response body
	})

	t.Run("TestUserEndpoint", func(t *testing.T) {
		w := performRequest(router, "GET", "/user")
		assert.Equal(t, http.StatusOK, w.Code)

		// TODO: Add assertions for the response body
	})
}

func setupRouter() *gin.Engine {
	router := adapter.NewRouter()
	server := adapter.NewServer(router)

	server.Routing()

	return server.Router.Engine
}

func performRequest(router *gin.Engine, method, path string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, path, nil)
	router.ServeHTTP(w, req)
	return w
}
