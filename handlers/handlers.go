package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/Wai30Yan/cna-server/driver"
	"github.com/Wai30Yan/cna-server/config"
	"github.com/Wai30Yan/cna-server/model"
	"github.com/Wai30Yan/cna-server/repository"
	"github.com/Wai30Yan/cna-server/repository/dbrepo"
)

var Repo *Repository 

type Repository struct {
	App *config.AppConfig
	DB repository.DatabaseRepo
}

func NewRepo(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: a,
		DB: dbrepo.NewPostgresRepo(db.SQL, a),
	}
}

func NewHandler(r *Repository) {
	Repo = r
}

func (m *Repository) GetAllSchedule(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hit endpoint: GetAllSchedule")
	schedules := m.DB.GetAllSchedules()

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(schedules)
}

func (m *Repository) GetScheduleByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hit endpoint: GetScheduleByID")
	sid := mux.Vars(r)["id"]
	id, err := strconv.Atoi(sid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	s := m.DB.GetScheduleByID(id)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s)
}

func (m *Repository) InsertSchedule(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hit endpoint: AddSchedule")
	var s model.Schedule

	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newS := m.DB.InsertSchedule(&s)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&newS)
}

func (m *Repository) UpdateSchedule(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hit endpoint: UpdateSchedule")

	sid := mux.Vars(r)["id"]
	id, err := strconv.Atoi(sid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var s model.Schedule

	err = json.NewDecoder(r.Body).Decode(&s)	
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	sche := m.DB.UpdateSchedule(id, &s)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&sche)
}

func (m *Repository) DeleteSchedule(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hit endpoint: DeleteSchedule")
	sid := mux.Vars(r)["id"]
	id, err := strconv.Atoi(sid)
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	res := m.DB.DeleteSchedule(id)

	fmt.Println(res)
}