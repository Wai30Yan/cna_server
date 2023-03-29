package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/Wai30Yan/cna-server/model"
)

func (m *Repository) GetAllClinics(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hit endpoint: GetAllClinics")
	clinics, err := m.DB.GetAllClinics()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&clinics)
}

func (m *Repository) GetClinicByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hit endpoint: GetClinicByID")
	sid := mux.Vars(r)["id"]
	id, err := strconv.Atoi(sid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	clinic, err := m.DB.GetClinicByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&clinic)
}

func (m *Repository) InsertClinic(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hit endpoint: InsertClinic")
	var clinic model.Clinic
	json.NewDecoder(r.Body).Decode(&clinic)
	newCln, err := m.DB.InsertClinic(&clinic)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(&newCln)
}

func (m *Repository) UpdateClinic(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hit endpoint: UpdateClinic")
	sid := mux.Vars(r)["id"]
	id, err := strconv.Atoi(sid)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	var clinic model.Clinic
	json.NewDecoder(r.Body).Decode(&clinic)

	updated, err := m.DB.UpdateClinic(id, clinic)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(&updated)
}

func (m *Repository) DeleteClinic(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hit endpoint: DeleteClinic")
	sid := mux.Vars(r)["id"]
	id, err := strconv.Atoi(sid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	res, err := m.DB.DeleteClinic(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(res)
}