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
	b := dorayaki.CreateDorayaki()
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
	if updateDorayaki.Rasa != "" {
		DorayakiDetails.Rasa = updateDorayaki.Rasa
	}
	if updateDorayaki.Deskripsi != "" {
		DorayakiDetails.Deskripsi = updateDorayaki.Deskripsi
	}
	if updateDorayaki.Gambar != "" {
		DorayakiDetails.Gambar = updateDorayaki.Gambar
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
