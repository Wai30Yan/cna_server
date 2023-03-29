package repository

import (
	"database/sql"

	"github.com/Wai30Yan/cna-server/pkg/model"
)

type DatabaseRepo interface {
	GetAllSchedules() []model.Schedule
	GetScheduleByID(id int) *model.Schedule
	InsertSchedule(s *model.Schedule) *model.Schedule
	UpdateSchedule(id int, updated *model.Schedule) *model.Schedule
	DeleteSchedule(id int) sql.Result

	// Doctor
	GetAllDoctors() ([]*model.Doctor, error)
	GetDoctorByID(id int) (*model.Doctor, error)
	InsertDoctor(doctor *model.Doctor) (*model.Doctor, error)
	UpdateDoctor(id int, doctor *model.Doctor) (*model.Doctor, error)
	DeleteDoctor(id int) (*sql.Result, error)

	// Appointment
	GetAllAppointments() ([]*model.Appointment, error)
	GetAppointmentByID(id int) (*model.Appointment, error)
	InsertAppointment(appointment *model.Appointment) (*model.Appointment, error)
	UpdateAppointment(id int, appointment model.Appointment) (*model.Appointment, error)
	DeleteAppointment(id int) (*sql.Result, error)
	
	// Clinics
	GetAllClinics() ([]*model.Clinic, error)
	GetClinicByID(id int) (*model.Clinic, error)
	InsertClinic(clinic *model.Clinic) (*model.Clinic, error)
	UpdateClinic(id int, clinic model.Clinic) (*model.Clinic, error)
	DeleteClinic(id int) (*sql.Result, error)
}