package test

import (
	"net/http"
	"testing"

	"gin-sample/router"

	"gopkg.in/go-playground/assert.v1"
)

func TestHealth(t *testing.T) {
	r := router.SetupRouter()
	w := performGet(r, "/health")
	assert.Equal(t, http.StatusOK, w.Code)
}
