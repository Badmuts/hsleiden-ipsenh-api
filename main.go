package main

import (
	"fmt"
	"time"

	"os"

	"github.com/badmuts/hsleiden-ipsenh-api/web"
	graceful "gopkg.in/tylerb/graceful.v1"
)

func main() {
	server := web.NewServer()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Starting server on: :" + port)
	graceful.Run(":"+port, 10*time.Second, server)
}
