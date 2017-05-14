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

// NewDatapointController creates the controller
func NewDatapointController(router *mux.Router, r *render.Render, db *mgo.Database) *DatapointController {
	ctrl := &DatapointController{router, r, db.C("datapoint"), db, model.DatapointModel(db)}
	ctrl.Register()
	return ctrl
}

// Register registers the routes with mux.Router
func (ctrl *DatapointController) Register() {
	ctrl.router.HandleFunc("/datapoints", ctrl.create).Name("datapoints.create").Methods("POST")
	ctrl.router.HandleFunc("/datapoints", ctrl.find).Name("datapoints.find").Methods("GET")
	ctrl.router.HandleFunc("/datapoints/{id}", ctrl.findOne).Name("datapoints.findOne").Methods("GET")
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

func (ctrl *DatapointController) findOne(res http.ResponseWriter, req *http.Request) {
	ID := mux.Vars(req)["id"]
	datapoint, err := ctrl.Datapoint.FindByID(ID)
	if err == mgo.ErrNotFound || err != nil {
		ctrl.r.JSON(res, http.StatusNotFound, datapoint)
		return
	}

	datapoint.Sensor, _ = datapoint.GetSensor()

	ctrl.r.JSON(res, http.StatusOK, datapoint)
}

func (ctrl *DatapointController) find(res http.ResponseWriter, req *http.Request) {
	datapoints, _ := ctrl.Datapoint.Find()

	for index, _ := range datapoints {
		datapoints[index].Sensor, _ = datapoints[index].GetSensor()
	}

	ctrl.r.JSON(res, http.StatusOK, datapoints)
}
