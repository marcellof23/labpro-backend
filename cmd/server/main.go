package main

import (
	"log"
	"net/http"

	"labpro-backend/pkg/config"
	"labpro-backend/pkg/routes"

	"github.com/gorilla/mux"
	_ "gorm.io/driver/mysql"
	_ "gorm.io/gorm"
)

func main() {
	r := mux.NewRouter()

	if err := config.ConnectDB(); err != nil {
		log.Fatal("Error when connecting to database")
	}
	log.Print("Success when connect to database")

	routes.DorayakiRoutesGroup(r)
	routes.DorayakiStoreRoutesGroup(r)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
