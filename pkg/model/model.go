package model

import (
	"time"
)

type Schedule struct {
	ID         int    `json:"id"`
	DoctorName string `json:"doctor_name"`
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
}

type Doctor struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Specialty string `json:"specialty"`
}

type Clinic struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

type Appointment struct {
	ID        int       `json:"id"`
	DoctorID  int       `json:"doctor_id"`
	ClinicID  int       `json:"clinic_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Doctor    Doctor    `json:"doctor"`
	Clinic    Clinic    `json:"clinic"`
}

type Availability struct {
	Monday    bool `json:"monday"`
	Tuesday   bool `json:"tuesday"`
	Wednesday bool `json:"wednesday"`
	Thursday  bool `json:"thursday"`
	Friday    bool `json:"friday"`
	Saturday  bool `json:"saturday"`
	Sunday    bool `json:"sunday"`
}
