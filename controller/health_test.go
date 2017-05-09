package controller

import (
	"testing"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

func TestHealthController(t *testing.T) {
	router := mux.NewRouter()
	r := render.New()

	NewHealthController(router, r)

	endPointRoute := router.Get("/healthz")
	if endPointRoute == nil {
		t.Errorf("/healthz endpoint is not registerd")
	}
}
