package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/badmuts/hsleiden-ipsenh-api/model"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// HubController represents the controller instance
type HubController struct {
	router *mux.Router
	r      *render.Render
	hubs   *mgo.Collection
	db     *mgo.Database
}

// NewHubController creates the controller
func NewHubController(router *mux.Router, r *render.Render, db *mgo.Database) *HubController {
	ctrl := &HubController{router, r, db.C("hub"), db}
	ctrl.Register()
	return ctrl
}

// Register registers the routes with mux.Router
func (c *HubController) Register() {
	c.router.HandleFunc("/hubs", c.create).Name("hubs.create").Methods("POST")
	c.router.HandleFunc("/hubs", c.find).Name("hubs.find").Methods("GET")
	c.router.HandleFunc("/hubs/{id}", c.findOne).Name("hubs.findOne").Methods("GET")
}

func (c *HubController) create(res http.ResponseWriter, req *http.Request) {
	var newHub model.Hub
	dec := json.NewDecoder(req.Body)
	dec.Decode(&newHub)

	newHub.ID = bson.NewObjectId()
	err := c.hubs.Insert(&newHub)

	if err != nil {
		c.r.JSON(res, http.StatusInternalServerError, err)
		log.Fatal(err)
	}

	c.r.JSON(res, http.StatusCreated, newHub)
}

func (c *HubController) findOne(res http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	hub := model.Hub{}

	err := c.hubs.FindId(bson.ObjectIdHex(id)).One(&hub)
	if err != nil {
		c.r.JSON(res, http.StatusInternalServerError, err)
		log.Fatal(err)
	}

	c.r.JSON(res, http.StatusOK, model.HubJSON{hub, hub.Sensors(c.db)})
}

func (c *HubController) find(res http.ResponseWriter, req *http.Request) {
	hubs := []model.Hub{}
	hubsJ := []model.HubJSON{}
	err := c.hubs.Find(bson.M{}).Limit(25).All(&hubs)

	if err != nil {
		c.r.JSON(res, http.StatusInternalServerError, err)
		log.Fatal(err)
	}

	// populate relationship
	for _, hub := range hubs {
		hubsJ = append(hubsJ, model.HubJSON{hub, hub.Sensors(c.db)})
	}

	c.r.JSON(res, http.StatusOK, hubsJ)
}
