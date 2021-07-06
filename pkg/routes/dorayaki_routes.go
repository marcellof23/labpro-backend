package routes

import (
	"labpro-backend/pkg/controllers"

	"github.com/gorilla/mux"
)

var DorayakiRoutesGroup = func(router *mux.Router) {
	router.HandleFunc("/dorayaki", controllers.CreateDorayaki).Methods("POST")
	router.HandleFunc("/dorayaki", controllers.GetDorayakis).Methods("GET")
	router.HandleFunc("/dorayaki/{dorayakiId}", controllers.GetDorayakiById).Methods("GET")
	router.HandleFunc("/dorayaki/{dorayakiId}", controllers.UpdateDorayaki).Methods("PUT")
	router.HandleFunc("/dorayaki/{dorayakiId}", controllers.DeleteDorayaki).Methods("DELETE")
}
