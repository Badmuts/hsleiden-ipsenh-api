package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"io/ioutil"

	"github.com/badmuts/hsleiden-ipsenh-api/model"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// HubController represents the controller instance
type SensorController struct {
	router  *mux.Router
	r       *render.Render
	sensors *mgo.Collection
	db      *mgo.Database
}

// NewSensorController creates the controller
func NewSensorController(router *mux.Router, r *render.Render, db *mgo.Database) *SensorController {
	ctrl := &SensorController{
		router:  router,
		r:       r,
		sensors: db.C("sensors"),
		db:      db,
	}
	ctrl.router.HandleFunc("/sensors/{id}/datapoints", ctrl.CreateDatapoints).Name("sensors.datapoint.create").Methods("POST")
	return ctrl
}

// CreateDatapoints saves datapoints to db
// Todo: this should be an alias method for the Create method in DatapointController
func (ctrl *SensorController) CreateDatapoints(res http.ResponseWriter, req *http.Request) {
	sensorID := mux.Vars(req)["id"]
	datapoints := make([]model.Datapoint, 0)
	dec, _ := ioutil.ReadAll(req.Body)
	err := json.Unmarshal(dec, &datapoints)

	log.Printf("sensorID: %s", sensorID)
	log.Printf("datapoints len: %s", len(datapoints))

	// save datapoints
	for index := range datapoints {
		datapoints[index].SensorID = bson.ObjectIdHex(sensorID)
	}

	datapoints, err = model.BulkSaveDatapoints(ctrl.db, datapoints)
	if err != nil {
		ctrl.r.JSON(res, http.StatusInternalServerError, err)
		log.Fatal(err)
	}

	ctrl.r.JSON(res, http.StatusCreated, datapoints)
}
