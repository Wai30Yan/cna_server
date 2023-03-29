package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/Wai30Yan/cna-server/model"
)

func (m *Repository) GetAllAppointments(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hit endpoint: GetAllAppointments")
	appointments, err := m.DB.GetAllAppointments()

	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&appointments)
}

func (m *Repository) InsertAppointment(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hit endpoint: InsertAppointment")
	var a model.Appointment

	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newA, err := m.DB.InsertAppointment(&a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Accept")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&newA)
}

func (m *Repository) GetAppointmentByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hit endpoint: GetAppointmentByID")
	sid := mux.Vars(r)["id"]
	id, err := strconv.Atoi(sid)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	a, err := m.DB.GetAppointmentByID(id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&a)
}

func (m *Repository) UpdateAppointment(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hit endpoint: UpdateAppointment")
	sid := mux.Vars(r)["id"]
	id, err := strconv.Atoi(sid)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	var a model.Appointment
	json.NewDecoder(r.Body).Decode(&a)

	updatedAppointment, err := m.DB.UpdateAppointment(id, a)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(&updatedAppointment)
}

func (m *Repository) DeleteAppointment(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hit endpoint: DeleteAppointment")
	sid := mux.Vars(r)["id"]
	id, err := strconv.Atoi(sid)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	res, err := m.DB.DeleteAppointment(id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	
	json.NewEncoder(w).Encode(res)
	// w.Write([]byte("Appointment successfully deleted."))
}