package routes

import (
	"labpro-backend/pkg/controllers"

	"github.com/gorilla/mux"
)

var DorayakiStoreRoutesGroup = func(router *mux.Router) {
	router.HandleFunc("/dorayaki-store/", controllers.CreateDorayakiStore).Methods("POST")
	router.HandleFunc("/dorayaki-store/", controllers.GetDorayakiStores).Methods("GET")
	router.HandleFunc("/dorayaki-store/{dorayakiStoreId}", controllers.GetDorayakiStoreById).Methods("GET")
	router.HandleFunc("/dorayaki-store/{dorayakiStoreId}", controllers.UpdateDorayakiStore).Methods("PUT")
	router.HandleFunc("/dorayaki-store/{dorayakiStoreId}", controllers.DeleteDorayakiStore).Methods("DELETE")
}
