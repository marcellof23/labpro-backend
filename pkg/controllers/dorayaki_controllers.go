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

var NewDorayaki models.Dorayaki

func CreateDorayaki(w http.ResponseWriter, r *http.Request) {
	dorayaki := &models.Dorayaki{}
	utils.ParseBody(r, dorayaki)
	b, err := dorayaki.CreateDorayaki()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		mapErr := map[string]string{"message": err.Error()}
		resError, _ := json.Marshal(mapErr)
		w.Write(resError)
		return
	}

	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetDorayakis(w http.ResponseWriter, r *http.Request) {
	newDorayakis := models.GetAllDorayakis()
	res, _ := json.Marshal(newDorayakis)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetDorayakiById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	DorayakiId := vars["DorayakiId"]
	ID, err := strconv.ParseInt(DorayakiId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	DorayakiDetails, _ := models.GetDorayakiById(ID)
	res, _ := json.Marshal(DorayakiDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetDorayakiByIdStore(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	DorayakiStoreId := vars["DorayakiIdStore"]
	ID, err := strconv.ParseInt(DorayakiStoreId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	DorayakiStoreDetails, _ := models.GetDorayakiByIdStore(ID)
	res, _ := json.Marshal(DorayakiStoreDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func VariantDorayaki(w http.ResponseWriter, r *http.Request) {
	VariantDorayakiStores := models.GetAllVariant()
	res, _ := json.Marshal(VariantDorayakiStores)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func VariantDorayakiById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	DorayakiId := vars["DorayakiId"]
	ID, err := strconv.ParseInt(DorayakiId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	DorayakiVariantsByID, _ := models.GetVariantById(ID)
	res, _ := json.Marshal(DorayakiVariantsByID)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateDorayaki(w http.ResponseWriter, r *http.Request) {
	var updateDorayaki = &models.Dorayaki{}
	utils.ParseBody(r, updateDorayaki)
	vars := mux.Vars(r)
	DorayakiId := vars["DorayakiId"]
	ID, err := strconv.ParseInt(DorayakiId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	DorayakiDetails, db := models.GetDorayakiById(ID)
	if updateDorayaki.Rasa != nil {
		DorayakiDetails.Rasa = updateDorayaki.Rasa
	}
	if updateDorayaki.Deskripsi != nil {
		DorayakiDetails.Deskripsi = updateDorayaki.Deskripsi
	}
	if updateDorayaki.Gambar != nil {
		DorayakiDetails.Gambar = updateDorayaki.Gambar
	}
	if updateDorayaki.DorayakiStoreID != nil {
		DorayakiDetails.DorayakiStoreID = updateDorayaki.DorayakiStoreID
	}
	db.Save(&DorayakiDetails)
	res, _ := json.Marshal(DorayakiDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteDorayaki(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	DorayakiId := vars["DorayakiId"]
	ID, err := strconv.ParseInt(DorayakiId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	Dorayaki := models.DeleteDorayaki(ID)
	res, _ := json.Marshal(Dorayaki)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
