package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/badmuts/hsleiden-ipsenh-api/model"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"gopkg.in/mgo.v2"
)

// HubController represents the controller instance
type HubController struct {
	router *mux.Router
	r      *render.Render
	hubs   *mgo.Collection
	db     *mgo.Database
	Hub    *model.Hub
}

// NewHubController creates the controller
func NewHubController(router *mux.Router, r *render.Render, db *mgo.Database) *HubController {
	ctrl := &HubController{router, r, db.C("hub"), db, model.HubModel(db)}
	ctrl.Register()
	return ctrl
}

// Register registers the routes with mux.Router
func (ctrl *HubController) Register() {
	ctrl.router.HandleFunc("/hubs", ctrl.create).Name("hubs.create").Methods("POST")
	ctrl.router.HandleFunc("/hubs", ctrl.find).Name("hubs.find").Methods("GET")
	ctrl.router.HandleFunc("/hubs/{id}", ctrl.findOne).Name("hubs.findOne").Methods("GET")
}

func (ctrl *HubController) create(res http.ResponseWriter, req *http.Request) {
	newHub := model.HubModel(ctrl.db)
	dec := json.NewDecoder(req.Body)
	err := dec.Decode(&newHub)

	_, err = newHub.Save(ctrl.db)

	if err != nil {
		ctrl.r.JSON(res, http.StatusInternalServerError, err)
		log.Fatal(err)
	}

	ctrl.r.JSON(res, http.StatusCreated, newHub)
}

func (ctrl *HubController) findOne(res http.ResponseWriter, req *http.Request) {
	ID := mux.Vars(req)["id"]
	hub, err := ctrl.Hub.FindByID(ctrl.db, ID)
	if err == mgo.ErrNotFound || err != nil {
		ctrl.r.JSON(res, http.StatusNotFound, hub)
		return
	}

	hJSON, err := hub.JSON(ctrl.db)
	if err != nil {
		ctrl.r.JSON(res, http.StatusInternalServerError, err)
		log.Fatal(err)
		return
	}

	ctrl.r.JSON(res, http.StatusOK, hJSON)
}

func (ctrl *HubController) find(res http.ResponseWriter, req *http.Request) {
	hubs, _ := ctrl.Hub.Find(ctrl.db)
	hubsJ := []model.HubJSON{}

	for _, hub := range hubs {
		hubJ, _ := hub.JSON(ctrl.db)
		hubsJ = append(hubsJ, hubJ)
	}

	ctrl.r.JSON(res, http.StatusOK, hubsJ)
}
