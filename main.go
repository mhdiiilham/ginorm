package main

import (
	db "github.com/mhdiiilham/ginorm/db"
	"github.com/mhdiiilham/ginorm/routers"
)

func main() {
	const PORT = ":8000"
	router := routers.Router()
	db.MySQL()
	router.Run(PORT)
}
