package main

import (
	"net/http"

	"github.com/Wai30Yan/cna-server/pkg/handlers"
	"github.com/gorilla/mux"
)

func routes() http.Handler {
	r := mux.NewRouter()

	admin := r.PathPrefix("/admin").Subrouter()
	admin.Use(authenticate)

	// gh.CORS(
	// 	gh.AllowedOrigins([]string{"*"}),
	// 	gh.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
	// )(r)
	// gh.CORS(
	// 	gh.AllowedOrigins([]string{"*"}),
	// 	gh.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
	// )(admin)

	r.Use(mux.CORSMethodMiddleware(r))

	// r.Use(preflightMiddleware)
	// admin.Use(preflightMiddleware)
	
	// Public HTTP GET Requests for doctors, clinics & appointments
	r.HandleFunc("/doctors", handlers.Repo.GetAllDoctors)
	r.HandleFunc("/doctor/{id}", handlers.Repo.GetDoctorByID)

	r.HandleFunc("/clinics", handlers.Repo.GetAllClinics)
	r.HandleFunc("/clinic/{id}", handlers.Repo.GetClinicByID)
	
	r.HandleFunc("/appointments", handlers.Repo.GetAllAppointments)
	r.HandleFunc("/appointment/{id}", handlers.Repo.GetAppointmentByID)
	
	// Admin Routes That Need Authentication for doctors, clinics & appointments
	admin.HandleFunc("/create-doctor", handlers.Repo.InsertDoctor).Methods("POST")
	admin.HandleFunc("/update-doctor/{id}", handlers.Repo.UpdateDoctor).Methods("PUT")
	admin.HandleFunc("/delete-doctor/{id}", handlers.Repo.DeleteDoctor).Methods("DELETE")
	
	admin.HandleFunc("/create-appointment", handlers.Repo.InsertAppointment).Methods("POST")
	admin.HandleFunc("/update-appointment/{id}", handlers.Repo.UpdateAppointment).Methods("PUT")
	admin.HandleFunc("/delete-appointment/{id}", handlers.Repo.DeleteAppointment).Methods("DELETE")

	admin.HandleFunc("/create-clinic", handlers.Repo.InsertClinic).Methods("POST")
	admin.HandleFunc("/update-clinic/{id}", handlers.Repo.UpdateClinic).Methods("PUT")
	admin.HandleFunc("/delete-clinic/{id}", handlers.Repo.DeleteClinic).Methods("DELETE")
	
	
	
	
	r.HandleFunc("/schedules", handlers.Repo.GetAllSchedule)
	r.HandleFunc("/schedule/{id}", handlers.Repo.GetScheduleByID)
	r.HandleFunc("/delete/{id}", handlers.Repo.DeleteSchedule).Methods("DELETE")
	r.HandleFunc("/update/{id}", handlers.Repo.UpdateSchedule).Methods("PUT")
	r.HandleFunc("/create", handlers.Repo.InsertSchedule).Methods("POST")

	return r
}