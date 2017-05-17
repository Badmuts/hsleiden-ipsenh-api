package controller

import (
	"encoding/json"
	"net/http"

	"errors"

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
	ctrl.router.HandleFunc("/buildings", ctrl.CreateBuilding).Name("buildings.create").Methods("POST")
	ctrl.router.HandleFunc("/buildings/{id}/rooms", ctrl.CreateRoom).Name("buildings.rooms.create").Methods("POST")
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
