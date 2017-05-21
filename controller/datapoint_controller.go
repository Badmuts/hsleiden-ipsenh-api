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
	body, err := ioutil.ReadAll(req.Body)
	er := json.Unmarshal(body, &newDatapoints)

	if er != nil {
		log.Fatal(er)
	}

	returnedDatapoints := []model.Datapoint{}
	occupationSensor := model.OccupationSensor{}
	for index := range newDatapoints {
		for i := range newDatapoints[index].Datapoints {
			newDatapoints[index].Datapoints[i].DB = ctrl.db
			newDatapoints[index].Datapoints[i].SensorID = newDatapoints[index].SensorID
			newDatapoints[index].Datapoints[i].Save()
			returnedDatapoints = append(returnedDatapoints, newDatapoints[index].Datapoints[i])

			hub := model.Hub{}
			err = ctrl.db.C("hub").Find(bson.M{"sensors": bson.M{"$elemMatch": bson.M{"_id": bson.ObjectIdHex("5915a9e7932c2b024d18561d")}}}).One(&hub)
			log.Printf("Hub: %s", hub)
			room := model.Room{}
			err = ctrl.db.C("room").Find(bson.M{"hubs": bson.M{"$elemMatch": bson.M{"_id": bson.ObjectIdHex("5915a9e7932c2b024d18561c")}}}).One(&room)
			log.Printf("Room: %s", room.MaxCapacity)

		}
	}

	if err != nil {
		ctrl.r.JSON(res, http.StatusInternalServerError, err)
		log.Fatal(err)
	}

	ctrl.r.JSON(res, http.StatusCreated, returnedDatapoints)
}
