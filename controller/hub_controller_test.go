package controller

import (
	"testing"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"gopkg.in/mgo.v2"
)

func TestHubController(t *testing.T) {
	router := mux.NewRouter()

	NewHubController(router, render.New(), &mgo.Database{})

	if router.Get("hubs.create") == nil {
		t.Errorf("No hubs.create endpoint registered")
	}

	if router.Get("hubs.find") == nil {
		t.Errorf("No hubs.find endpoint registered")
	}

	if router.Get("hubs.findOne") == nil {
		t.Errorf("No hubs.findOne endpoint registered")
	}

	if router.Get("hubs.findOne") == nil {
		t.Errorf("No hubs.findOne endpoint registered")
	}
}
