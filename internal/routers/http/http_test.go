package http_test

import (
	r "address-book/internal/routers/http"

	"github.com/stretchr/testify/assert"

	"net/http"
	"net/http/httptest"
	"testing"
)

const ConfigPath = "/config/config.yml"

var (
	router = r.SetupRouter()
	w      = httptest.NewRecorder()
)

func TestPingRoute(t *testing.T) {
	req, _ := http.NewRequest(
		"GET",
		"/ping",
		nil,
	)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
