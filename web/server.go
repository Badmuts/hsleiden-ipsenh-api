package web

import (
	"github.com/badmuts/hsleiden-ipsenh-api/controller"
	"github.com/badmuts/hsleiden-ipsenh-api/db"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

// Server represents a server based on Negroni
type Server struct {
	*negroni.Negroni
}

// NewServer creates a new server
func NewServer() *Server {
	r := render.New()
	router := mux.NewRouter()
	db := db.Connect()
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})

	controller.NewHealthController(router, r)
	controller.NewHubController(router, r, db)
	controller.NewDatapointController(router, r, db)
	controller.NewSensorController(router, r, db)
	controller.NewBuildingController(router, r, db)

	server := Server{negroni.Classic()}
	server.Use(c)
	server.UseHandler(router)

	return &server
}
