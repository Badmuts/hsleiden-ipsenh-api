package controller

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/badmuts/hsleiden-ipsenh-api/model"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type RoomController struct {
	router *mux.Router
	r      *render.Render
	db     *mgo.Database
}

func NewRoomController(router *mux.Router, r *render.Render, db *mgo.Database) *RoomController {
	ctrl := &RoomController{
		router: router,
		r:      r,
		db:     db,
	}
	ctrl.Register()
	return ctrl
}

func (ctrl *RoomController) Register() {
	ctrl.router.HandleFunc("/rooms", ctrl.Find).Name("rooms.find").Methods("GET")
	ctrl.router.HandleFunc("/rooms", ctrl.Create).Name("rooms.create").Methods("POST")
	ctrl.router.HandleFunc("/rooms/{roomID}", ctrl.FindRoomID).Name("rooms.findId").Methods("GET")
	ctrl.router.HandleFunc("/rooms/{roomID}", ctrl.Update).Name("rooms.update").Methods("PUT", "PATCH")
	ctrl.router.HandleFunc("/rooms/{roomID}/roster", ctrl.CreateRoster).Name("rooms.createRoster").Methods("POST")
}

func (ctrl *RoomController) FindRoomID(res http.ResponseWriter, req *http.Request) {
	roomID := bson.ObjectIdHex(mux.Vars(req)["roomID"])
	Room := model.NewRoomModel(ctrl.db)
	room, err := Room.FindId(roomID)
	if err == mgo.ErrNotFound {
		ctrl.r.JSON(res, http.StatusNotFound, NewControllerError("404", "Can't find room due to db error (enable mgo logging)", "Could not find room", "more info"))
		return
	} else if err != nil {
		ctrl.r.JSON(res, http.StatusInternalServerError, NewControllerError("500", "Can't find room due to db error (enable mgo logging)", "Can't find room", "more info"))
		return
	}

	ctrl.r.JSON(res, http.StatusOK, room)
}

func (ctrl *RoomController) Update(res http.ResponseWriter, req *http.Request) {
	roomID := bson.ObjectIdHex(mux.Vars(req)["roomID"])
	Room := model.NewRoomModel(ctrl.db)
	Room, err := Room.FindId(roomID)
	if err == mgo.ErrNotFound {
		ctrl.r.JSON(res, http.StatusNotFound, Room)
		return
	} else if err != nil {
		ctrl.r.JSON(res, http.StatusInternalServerError, NewControllerError("500", "Can't update due to db error (enable mgo logging)", "Can't update room", "more info"))
		return
	}

	dec := json.NewDecoder(req.Body)
	err = dec.Decode(&Room)
	if err != nil {
		ctrl.r.JSON(res, http.StatusBadRequest, NewControllerError("400", "Can't decode due to invalid json", "Invalid json", "more info"))
		return
	}

	Room, err = Room.Save()
	if err != nil {
		ctrl.r.JSON(res, http.StatusInternalServerError, NewControllerError("500", "Can't UpSert due to db error (enable mgo logging)", "Could not save room", "more info"))
		return
	}

	ctrl.r.JSON(res, http.StatusOK, Room)
}

func (ctrl *RoomController) CreateRoster(res http.ResponseWriter, req *http.Request) {
	ID := bson.ObjectIdHex(mux.Vars(req)["roomID"])
	RoomRosters := []model.RoomRoster{}
	dec := json.NewDecoder(req.Body)
	err := dec.Decode(&RoomRosters)
	if err != nil {
		ctrl.r.JSON(res, http.StatusBadRequest, NewControllerError("400", "Can't decode due to invalid json", "Invalid json", "more info"))
		return
	}

	for index := range RoomRosters {
		if RoomRosters[index].ID == "" {
			RoomRosters[index].ID = bson.NewObjectId()
		}
		RoomRosters[index].RoomID = ID

		err = ctrl.db.C("room_reservation").Insert(RoomRosters[index])
		if err != nil {
			ctrl.r.JSON(res, http.StatusInternalServerError, NewControllerError("500", "Can't insert due to db error (enable mgo logging)", "Could not save reservation", "more info"))
			return
		}
	}

	ctrl.r.JSON(res, http.StatusCreated, RoomRosters)
}

func (ctrl *RoomController) Find(res http.ResponseWriter, req *http.Request) {
	building_id := req.URL.Query().Get("buildingId")
	room := model.NewRoomModel(ctrl.db)
	rooms, err := room.Find(building_id)
	if err != nil {
		ctrl.r.JSON(res, http.StatusInternalServerError, NewControllerError("500", "Can't find due to db error (enable mgo logging)", "Can't find room", "more info"))
		return
	}

	ctrl.r.JSON(res, http.StatusOK, rooms)
}

func (ctrl *RoomController) Create(res http.ResponseWriter, req *http.Request) {
	Room := model.NewRoomModel(ctrl.db)

	dec := json.NewDecoder(req.Body)
	err := dec.Decode(&Room)
	if err != nil {
		ctrl.r.JSON(res, http.StatusBadRequest, NewControllerError("400", "Can't decode due to invalid json", "Invalid json", "more info"))
		return
	}

	Room, err = Room.Save()
	if err != nil {
		ctrl.r.JSON(res, http.StatusInternalServerError, errors.New("Could not save room"))
		return
	}

	ctrl.r.JSON(res, http.StatusCreated, Room)
}
