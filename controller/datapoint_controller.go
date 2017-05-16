package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/badmuts/hsleiden-ipsenh-api/model"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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
	SensorID   bson.ObjectId     `json:"sensor_id"`
	Datapoints []model.Datapoint `json:"data"`
	db         *mgo.Database
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
	newDatapoints := []datapoints{}
	// dec := json.NewDecoder(req.Body)
	// err := dec.Decode(&newDatapoints)
	body, err := ioutil.ReadAll(req.Body)
	er := json.Unmarshal(body, &newDatapoints)
	if er != nil {
		panic(er)
	}

	// log.Printf("datapoints data %s", newDatapoints.Datapoints)
	for index := range newDatapoints {
		for i := range newDatapoints[index].Datapoints {
			dp := *model.DatapointModel(ctrl.db)
			dp.SensorID = newDatapoints[index].SensorID
			dp.Key = newDatapoints[index].Datapoints[i].Key
			dp.Value = newDatapoints[index].Datapoints[i].Value
			dp.Timestamp = newDatapoints[index].Datapoints[i].Timestamp
			_, err = dp.Save()
		}
	}
	// _, err = newDatapoint.Save()

	if err != nil {
		ctrl.r.JSON(res, http.StatusInternalServerError, err)
		log.Fatal(err)
	}

	ctrl.r.JSON(res, http.StatusCreated, newDatapoints)
}
