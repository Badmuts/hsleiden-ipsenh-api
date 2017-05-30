package controller

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"time"

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
	SensorID   bson.ObjectId      `json:"sensor_id"`
	SensorType string             `json:"name"`
	Datapoints []*model.Datapoint `json:"data"`
}

type RoomLog struct {
	ID         bson.ObjectId `json:"-" bson:"_id"`
	RoomID     bson.ObjectId `json:"-" bson:"room"`
	Occupation int           `json:"-" bson:"occupation"`
	Timestamp  time.Time     `json:"-" bson:"timestamp"`
}

type RoomLog struct {
	ID         bson.ObjectId `json:"-" bson:"_id"`
	RoomID     bson.ObjectId `json:"-" bson:"room"`
	Occupation int           `json:"-" bson:"occupation"`
	Timestamp  time.Time     `json:"-" bson:"timestamp"`
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
	hub := model.Hub{}
	room := model.Room{}
	for index := range newDatapoints {
		for i := range newDatapoints[index].Datapoints {
			newDatapoints[index].Datapoints[i].DB = ctrl.db
			newDatapoints[index].Datapoints[i].SensorID = newDatapoints[index].SensorID
			newDatapoints[index].Datapoints[i].Save()
			returnedDatapoints = append(returnedDatapoints, *newDatapoints[index].Datapoints[i])

			err = ctrl.db.C("hub").Find(bson.M{"sensors": newDatapoints[index].SensorID}).One(&hub)
			if err != nil {
				log.Printf("HUB ERROR %s", err)
				ctrl.r.JSON(res, http.StatusInternalServerError, errors.New("Could not find hub of with sensor_id"))
				return
			}

			err = ctrl.db.C("room").Find(bson.M{"hubs": hub.ID}).One(&room)
			if err != nil {
				log.Printf("ROOM ERROR %s", err)
				ctrl.r.JSON(res, http.StatusInternalServerError, errors.New("Could not find room of hub"))
				return
			}

			if room.ID != "" {
				if newDatapoints[index].SensorType == "in" {
					occupationSensor.InSensorDatapoints = append(occupationSensor.InSensorDatapoints, newDatapoints[index].Datapoints[i])
				}
				if newDatapoints[index].SensorType == "out" {
					occupationSensor.OutSensorDatapoints = append(occupationSensor.OutSensorDatapoints, newDatapoints[index].Datapoints[i])
				}
			}
		}
	}

	if room.ID != "" {
		ctrl.CalculateAndUpdateRoomOccupation(occupationSensor, room)
	}

	if err != nil {
		ctrl.r.JSON(res, http.StatusInternalServerError, err)
		log.Fatal(err)
	}

	ctrl.r.JSON(res, http.StatusCreated, returnedDatapoints)
}

func (ctrl *DatapointController) CalculateAndUpdateRoomOccupation(o model.OccupationSensor, r model.Room) (info *mgo.ChangeInfo, err error) {
	entrances := o.CalculateEntrances()
	log.Printf("entrances: %s", entrances)

	exits := o.CalculateExits()
	log.Printf("exits: %s", exits)

	tmpOccupation := (r.Occupation - exits)

	r.Occupation = tmpOccupation + entrances
	log.Printf("occupation %s", r.Occupation)

	log.Printf("Room %s", r)
	// log.Printf("Database %s", ctrl.db)
	// info, err = ctrl.db.C("room").UpsertId(r.ID, r)

	// if err != nil {
	// 	return info, err
	// }

	roomLog := &RoomLog{}
	roomLog.ID = bson.NewObjectId()
	roomLog.RoomID = r.ID
	roomLog.Occupation = r.Occupation
	roomLog.Timestamp = time.Now()

	_, err = ctrl.db.C("room_log").Upsert(roomLog, &roomLog)

	if err != nil {
		return nil, err
	}

	return info, err
}
