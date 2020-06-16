package main

import (
	db "github.com/mhdiiilham/ginorm/db"
	"net/http"
	"github.com/mhdiiilham/ginorm/routers"
	log "github.com/sirupsen/logrus"
)

func main() {
	port := ":8000"
	router := routers.Router()
	db.MySQL()
	log.Fatal(http.ListenAndServe(port, router))
}