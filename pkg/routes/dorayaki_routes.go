package routes

import (
	"labpro-backend/pkg/controllers"

	"github.com/gorilla/mux"
)

var DorayakiRoutesGroup = func(router *mux.Router) {
	router.HandleFunc("/dorayaki/store/{DorayakiIdStore}", controllers.GetDorayakiByIdStore).Methods("GET")
	router.HandleFunc("/dorayaki/store/{DorayakiIdStore}", controllers.UpdateDorayakiByIdStore).Methods("PUT")
	router.HandleFunc("/dorayaki/transfer", controllers.UpdateDorayakiByIdStore).Methods("POST")
	router.HandleFunc("/dorayaki/variant/{DorayakiId}", controllers.VariantDorayakiById).Methods("GET")
	router.HandleFunc("/dorayaki/variant", controllers.VariantDorayaki).Methods("GET")
	router.HandleFunc("/dorayaki/{DorayakiId}", controllers.UpdateDorayaki).Methods("PUT")
	router.HandleFunc("/dorayaki/{DorayakiId}", controllers.DeleteDorayaki).Methods("DELETE")
	router.HandleFunc("/dorayaki/{DorayakiId}", controllers.GetDorayakiById).Methods("GET")
	router.HandleFunc("/dorayaki", controllers.CreateDorayaki).Methods("POST")
	router.HandleFunc("/dorayaki", controllers.GetDorayakis).Methods("GET")
}
