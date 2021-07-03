package controllers

import (
	"encoding/json"
	"fmt"
	"labpro-backend/pkg/models"
	"labpro-backend/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewDorayakiStore models.DorayakiStore

func CreateDorayakiStore(w http.ResponseWriter, r *http.Request) {
	DorayakiStore := &models.DorayakiStore{}
	utils.ParseBody(r, DorayakiStore)
	b := DorayakiStore.CreateDorayakiStore()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetDorayakiStores(w http.ResponseWriter, r *http.Request) {
	newDorayakiStores := models.GetAllDorayakiStores()
	res, _ := json.Marshal(newDorayakiStores)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetDorayakiStoreById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	DorayakiStoreId := vars["DorayakiStoreId"]
	ID, err := strconv.ParseInt(DorayakiStoreId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	DorayakiStoreDetails, _ := models.GetDorayakiStoreById(ID)
	res, _ := json.Marshal(DorayakiStoreDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateDorayakiStore(w http.ResponseWriter, r *http.Request) {
	var updateDorayakiStore = &models.DorayakiStore{}
	utils.ParseBody(r, updateDorayakiStore)
	vars := mux.Vars(r)
	DorayakiStoreId := vars["DorayakiStoreId"]
	ID, err := strconv.ParseInt(DorayakiStoreId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	DorayakiStoreDetails, db := models.GetDorayakiStoreById(ID)
	if updateDorayakiStore.Nama != "" {
		DorayakiStoreDetails.Nama = updateDorayakiStore.Nama
	}
	if updateDorayakiStore.Jalan != "" {
		DorayakiStoreDetails.Jalan = updateDorayakiStore.Jalan
	}
	if updateDorayakiStore.Provinsi != "" {
		DorayakiStoreDetails.Provinsi = updateDorayakiStore.Provinsi
	}
	if updateDorayakiStore.Kecamatan != "" {
		DorayakiStoreDetails.Kecamatan = updateDorayakiStore.Kecamatan
	}

	db.Save(&DorayakiStoreDetails)

	res, _ := json.Marshal(DorayakiStoreDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteDorayakiStore(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	DorayakiStoreId := vars["DorayakiStoreId"]
	ID, err := strconv.ParseInt(DorayakiStoreId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	DorayakiStore := models.DeleteDorayakiStore(ID)
	res, _ := json.Marshal(DorayakiStore)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
