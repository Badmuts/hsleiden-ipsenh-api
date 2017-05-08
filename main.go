package main

import (
	"fmt"
	"net/http"
	"time"

	"os"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
	graceful "gopkg.in/tylerb/graceful.v1"
)

func main() {
	r := render.New()
	router := mux.NewRouter()

	router.HandleFunc("/healthz", func(w http.ResponseWriter, req *http.Request) {
		r.JSON(w, http.StatusOK, map[string]bool{"healthy": true})
	})

	server := negroni.Classic()
	server.UseHandler(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Starting server on: :" + port)
	graceful.Run(":"+port, 10*time.Second, server)
}
