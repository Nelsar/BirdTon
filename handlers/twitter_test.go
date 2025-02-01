package handlers_test

import (
	"birdton/configurations"
	"birdton/handlers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kkdai/twitter"
	"github.com/stretchr/testify/assert"
)

func TestRedirectUserToTwitter(t *testing.T) {

	config := configurations.GetConfiguration()

	xClient := twitter.NewServerClient(config.ConsumerKey, config.ConsumerSecret)

	client := handlers.NewServerClient(xClient, config)

	req, err := http.NewRequest("GET", "/request", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(client.RedirectUserToTwitter)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}
