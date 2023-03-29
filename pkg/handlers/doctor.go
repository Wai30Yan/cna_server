package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/Wai30Yan/cna-server/pkg/model"
)

func (m *Repository) GetAllDoctors(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hit endpoint: GetAllDoctors")
	doctors, err := m.DB.GetAllDoctors()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(&doctors)
}

func (m *Repository) GetDoctorByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hit endpoint: GetDoctorByID")
	sid := mux.Vars(r)["id"]
	id, err := strconv.Atoi(sid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	doctor, err := m.DB.GetDoctorByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Accept")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&doctor)
}

func (m *Repository) InsertDoctor(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hit endpoint: InsertDoctor")
	var d model.Doctor

	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newDoc, err := m.DB.InsertDoctor(&d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Accept")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&newDoc)
}

func (m *Repository) UpdateDoctor(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hit endpoint: UpdateDoctor")
	sid := mux.Vars(r)["id"]
	var doctor *model.Doctor
	json.NewDecoder(r.Body).Decode(&doctor)
	id, err := strconv.Atoi(sid)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	updatedDoc, err := m.DB.UpdateDoctor(id, doctor)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(&updatedDoc)
}
func (m *Repository) DeleteDoctor(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hit endpoint: DeleteDoctor")
	sid := mux.Vars(r)["id"]
	id, err := strconv.Atoi(sid)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	res, err := m.DB.DeleteDoctor(id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Accept")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}