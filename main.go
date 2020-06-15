package main

import (
	"net/http"
	"github.com/mhdiiilham/ginorm/routers"
	log "github.com/sirupsen/logrus"
)

func main() {
	port := ":8000"
	router := routers.Router()
	log.Fatal(http.ListenAndServe(port, router))
}