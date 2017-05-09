package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// HealthController represents the controller instance
type HealthController struct {
	router *mux.Router
	r      *render.Render
}

// NewHealthController creates the controller
func NewHealthController(router *mux.Router, r *render.Render) *HealthController {
	ctrl := &HealthController{router, r}
	ctrl.Register()
	return ctrl
}

// Register registers the routes with mux.Router
func (c *HealthController) Register() {
	c.router.HandleFunc("/healthz", c.healthz).Name("/healthz")
}

func (c *HealthController) healthz(res http.ResponseWriter, req *http.Request) {
	c.r.JSON(res, http.StatusOK, map[string]bool{"healthy": true})
}
