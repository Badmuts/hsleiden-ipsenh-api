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
}
