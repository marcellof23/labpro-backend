package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"labpro-backend/pkg/config"
	"labpro-backend/pkg/routes"
	"labpro-backend/pkg/seeds"

	"github.com/go-chi/cors"
	"github.com/gorilla/mux"
	_ "gorm.io/driver/mysql"
	_ "gorm.io/gorm"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func seeding(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You have seeds all the models!")
	fmt.Println("Endpoint Hit: seeding")

	for _, seed := range seeds.AllDorayakiStore() {
		if err := seed.Run(config.DB); err != nil {
			log.Fatalf("Running seed '%s', failed with error: %s", seed.Name, err)
		}
	}

	for _, seed := range seeds.AllDorayaki() {
		if err := seed.Run(config.DB); err != nil {
			log.Fatalf("Running seed '%s', failed with error: %s", seed.Name, err)
		}
	}

}

func main() {
	r := mux.NewRouter()

	err := config.ConnectDB()
	if err != nil {
		log.Fatal("Error when connecting to database")
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = config.DefaultPort
	}
	// config.DB.AutoMigrate(&models.Dorayaki{}, &models.DorayakiStore{})

	log.Println("Success when connect to database")

	routes.DorayakiRoutesGroup(r)
	routes.DorayakiStoreRoutesGroup(r)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedHeaders:   []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Access-Control-Allow-Origin", "Accept"},
		AllowedMethods:   []string{"*"},
		AllowCredentials: true,
	})

	r.HandleFunc("/homepage", homePage)
	r.HandleFunc("/seeds", seeding)
	handler := c.Handler(r)

	log.Println("Connected to port http://localhost:8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", handler))

}
