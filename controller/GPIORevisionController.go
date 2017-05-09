package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/badmuts/hsleiden-ipsenh-api/model"
)

// GPIORevisionController represents the controller instance
type GPIORevisionController struct {
	router *mux.Router
	r      *render.Render
}

// NewGPIORevisionController creates the controller
func NewGPIORevisionController(router *mux.Router, r *render.Render) *GPIORevisionController {
	ctrl := &GPIORevisionController{router, r}
	ctrl.Register()
	return ctrl
}

// Register registers the routes with mux.Router
func (c *GPIORevisionController) Register() {
	c.router.HandleFunc("/gpiorevision", c.getRevision).Methods("POST")

}

func (c *GPIORevisionController) getRevision(res http.ResponseWriter, req *http.Request) {
	a := model.GPIORevision{Name: "Model B+", RevisionCode: "0010", Pins: 17}

	c.r.JSON(res, http.StatusOK, a)
}
