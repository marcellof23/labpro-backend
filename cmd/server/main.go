package main

import (
	"fmt"
	"log"
	"net/http"

	"labpro-backend/pkg/config"
	"labpro-backend/pkg/routes"
	"labpro-backend/pkg/seeds"

	"github.com/gorilla/mux"
	_ "gorm.io/driver/mysql"
	_ "gorm.io/gorm"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func main() {
	r := mux.NewRouter()

	err := config.ConnectDB()
	if err != nil {
		log.Fatal("Error when connecting to database")
	}

	for _, seed := range seeds.All1() {
		if err := seed.Run(config.DB); err != nil {
			log.Fatalf("Running seed '%s', failed with error: %s", seed.Name, err)
		}
	}

	for _, seed := range seeds.All2() {
		if err := seed.Run(config.DB); err != nil {
			log.Fatalf("Running seed '%s', failed with error: %s", seed.Name, err)
		}
	}

	log.Println("Success when connect to database")

	routes.DorayakiRoutesGroup(r)
	routes.DorayakiStoreRoutesGroup(r)

	r.HandleFunc("/tes", homePage)
	http.Handle("/", r)

	log.Println("Connected to port http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
