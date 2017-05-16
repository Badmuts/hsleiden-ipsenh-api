package controller

import (
	"testing"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"gopkg.in/mgo.v2"
)

func TestDatapointController(t *testing.T) {
	router := mux.NewRouter()

	NewDatapointController(router, render.New(), &mgo.Database{})

	if router.Get("datapoints.create") == nil {
		t.Errorf("No datapoints.create endpoint registered")
	}

	if router.Get("datapoints.find") == nil {
		t.Errorf("No datapoints.find endpoint registered")
	}

	if router.Get("datapoints.findOne") == nil {
		t.Errorf("No datapoints.findOne endpoint registered")
	}

	if router.Get("datapoints.findOne") == nil {
		t.Errorf("No datapoints.findOne endpoint registered")
	}
}
