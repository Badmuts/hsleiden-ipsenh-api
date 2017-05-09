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

// RegisterController represents the controller instance
type RegisterController struct {
	router *mux.Router
	r      *render.Render
	db     *mgo.Database
}

// NewRegisterController creates the controller
func NewRegisterController(router *mux.Router, r *render.Render, db *mgo.Database) *RegisterController {
	ctrl := &RegisterController{router, r, db}
	ctrl.Register()
	return ctrl
}

// Register registers the routes with mux.Router
func (c *RegisterController) Register() {
	c.router.HandleFunc("/register", c.register).Name("/register").Methods("POST")
}

func (c *RegisterController) register(res http.ResponseWriter, req *http.Request) {
	var newHub model.Hub
	dec := json.NewDecoder(req.Body)
	dec.Decode(&newHub)

	col := c.db.C("hub")
	err := col.Insert(&newHub)

	if err != nil {
		c.r.JSON(res, http.StatusInternalServerError, err)
		log.Fatal(err)
	}

	c.r.JSON(res, http.StatusCreated, newHub)
}
