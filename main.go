package main

import (
	db "github.com/mhdiiilham/ginorm/db"
	"github.com/mhdiiilham/ginorm/routers"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	port := ":8000"
	router := routers.Router()
	db.MySQL()
	log.Fatal(http.ListenAndServe(port, router))
}
