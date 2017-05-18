package controller

import (
	"encoding/json"
	"net/http"

	"errors"

	"log"

	"github.com/badmuts/hsleiden-ipsenh-api/model"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type BuildingController struct {
	router *mux.Router
	r      *render.Render
	db     *mgo.Database
}

func NewBuildingController(router *mux.Router, r *render.Render, db *mgo.Database) *BuildingController {
	ctrl := &BuildingController{
		router: router,
		r:      r,
		db:     db,
	}
	ctrl.Register()
	return ctrl
}

func (ctrl *BuildingController) Register() {
	ctrl.router.HandleFunc("/buildings", ctrl.FindBuilding).Name("buildings.find").Methods("GET")
	ctrl.router.HandleFunc("/buildings", ctrl.CreateBuilding).Name("buildings.create").Methods("POST")
	ctrl.router.HandleFunc("/buildings/{id}/rooms", ctrl.CreateRoom).Name("buildings.rooms.create").Methods("POST")
	ctrl.router.HandleFunc("/buildings/{id}/rooms", ctrl.FindRooms).Name("buildings.rooms.find").Methods("GET")
}

func (ctrl *BuildingController) FindBuilding(res http.ResponseWriter, req *http.Request) {
	building := model.NewBuildingModel(ctrl.db)
	buildings, err := building.Find()
	if err != nil {
		log.Fatal("could not retrieve buildings: ", err)
		ctrl.r.JSON(res, http.StatusInternalServerError, err)
		return
	}

	ctrl.r.JSON(res, http.StatusOK, buildings)
}

func (ctrl *BuildingController) CreateBuilding(res http.ResponseWriter, req *http.Request) {
	Building := model.NewBuildingModel(ctrl.db)

	dec := json.NewDecoder(req.Body)
	err := dec.Decode(&Building)
	if err != nil {
		ctrl.r.JSON(res, http.StatusInternalServerError, errors.New("Invalid json"))
		return
	}

	Building, err = Building.Save()
	if err != nil {
		ctrl.r.JSON(res, http.StatusInternalServerError, errors.New("Could not save building"))
	}

	ctrl.r.JSON(res, http.StatusCreated, Building)
}

// TODO move to RoomCtrl
func (ctrl *BuildingController) CreateRoom(res http.ResponseWriter, req *http.Request) {
	ID := mux.Vars(req)["id"]
	Room := model.NewRoomModel(ctrl.db)

	dec := json.NewDecoder(req.Body)
	err := dec.Decode(&Room)
	if err != nil {
		ctrl.r.JSON(res, http.StatusInternalServerError, errors.New("Invalid json"))
		return
	}

	Room.BuildingID = bson.ObjectIdHex(ID)

	Room, err = Room.Save()
	if err != nil {
		ctrl.r.JSON(res, http.StatusInternalServerError, errors.New("Could not save room"))
		return
	}

	ctrl.r.JSON(res, http.StatusCreated, Room)
}

// TODO move to RoomCtrl
func (ctrl *BuildingController) FindRooms(res http.ResponseWriter, req *http.Request) {
	buildingID := bson.ObjectIdHex(mux.Vars(req)["id"])
	building := model.NewBuildingModel(ctrl.db)
	building = building.FindId(buildingID)

	rooms, err := building.GetRooms()
	if err != nil {
		log.Fatal("Could not find rooms", err)
		ctrl.r.JSON(res, http.StatusNotFound, rooms)
		return
	}

	ctrl.r.JSON(res, http.StatusOK, rooms)
}
