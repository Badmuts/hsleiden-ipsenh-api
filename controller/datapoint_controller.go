package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/badmuts/hsleiden-ipsenh-api/model"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	mgo "gopkg.in/mgo.v2"
)

// DatapointController represents the controller instance
type DatapointController struct {
	router     *mux.Router
	r          *render.Render
	datapoints *mgo.Collection
	db         *mgo.Database
	Datapoint  *model.Datapoint
}

type datapoints struct {
	Sensor     *model.Sensor
	Datapoints []model.Datapoint
}

// NewDatapointController creates the controller
func NewDatapointController(router *mux.Router, r *render.Render, db *mgo.Database) *DatapointController {
	ctrl := &DatapointController{router, r, db.C("datapoint"), db, model.DatapointModel(db)}
	ctrl.Register()
	return ctrl
}

// Register registers the routes with mux.Router
func (ctrl *DatapointController) Register() {
	ctrl.router.HandleFunc("/datapoints", ctrl.create).Name("datapoints.create").Methods("POST")
}

func (ctrl *DatapointController) create(res http.ResponseWriter, req *http.Request) {
	newDatapoint := model.DatapointModel(ctrl.db)
	dec := json.NewDecoder(req.Body)
	err := dec.Decode(&newDatapoint)

	_, err = newDatapoint.Save()

	if err != nil {
		ctrl.r.JSON(res, http.StatusInternalServerError, err)
		log.Fatal(err)
	}

	ctrl.r.JSON(res, http.StatusCreated, newDatapoint)
}
