package dbrepo

import (
	"database/sql"
	"fmt"

	"github.com/Wai30Yan/cna-server/pkg/model"
)

func (m *postgresDBRepo) GetAllAppointments() ([]*model.Appointment, error) {
	var appointments []*model.Appointment 
	query := `select a.id, a.doctor_id, a.clinic_id, a.start_time, a.end_time, d.id, d.name, d.specialty, c.id, c.cname, c.location
	          from appointments a
			  join doctors d on a.doctor_id = d.id
			  join clinics c on a.clinic_id = c.id
	`

	rows, err := m.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var a model.Appointment
		err := rows.Scan(&a.ID, &a.DoctorID, &a.ClinicID, &a.StartTime, &a.EndTime, 
			&a.Doctor.ID, &a.Doctor.Name, &a.Doctor.Specialty, 
			&a.Clinic.ID, &a.Clinic.Name, &a.Clinic.Location)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		appointments = append(appointments, &a)
	}
	return appointments, nil
}


func (m *postgresDBRepo) GetAppointmentByID(id int) (*model.Appointment, error) {
	query := `select * from appointments a join doctors d on a.doctor_id = d.id
	join clinics c on a.clinic_id = c.id where a.id = $1
	`
	var a model.Appointment
	err := m.DB.QueryRow(query, id).Scan(&a.ID, &a.DoctorID, &a.ClinicID, &a.StartTime, &a.EndTime,
		&a.Doctor.ID, &a.Doctor.Name, &a.Doctor.Specialty,
		&a.Clinic.ID, &a.Clinic.Name, &a.Clinic.Location)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &a, nil
}

func (m *postgresDBRepo) InsertAppointment(a *model.Appointment) (*model.Appointment, error) {
	query := `WITH new_appointment AS (
		INSERT INTO appointments (doctor_id, clinic_id, start_time, end_time)
		VALUES ($1,$2,$3,$4)
		RETURNING id, doctor_id, clinic_id, start_time, end_time
	  )
	  SELECT
		new_appointment.id,
		new_appointment.doctor_id,
		new_appointment.clinic_id,
		new_appointment.start_time,
		new_appointment.end_time,
		doctors.id AS id,
		doctors.name AS doctor_name,
		doctors.specialty AS doctor_specialty,
		clinics.id AS id,
		clinics.cname AS clinic_name,
		clinics.location AS clinic_location
	  FROM new_appointment
	  JOIN doctors ON new_appointment.doctor_id = doctors.id
	  JOIN clinics ON new_appointment.clinic_id = clinics.id;`


	row := m.DB.QueryRow(query, a.DoctorID, a.ClinicID, a.StartTime, a.EndTime)

	err := row.Scan(&a.ID, &a.DoctorID, &a.ClinicID, &a.StartTime, &a.EndTime,
		&a.Doctor.ID, &a.Doctor.Name, &a.Doctor.Specialty,
		&a.Clinic.ID, &a.Clinic.Name, &a.Clinic.Location)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return a, nil
}

func (m *postgresDBRepo) UpdateAppointment(id int, a model.Appointment) (*model.Appointment, error) {
	query := `with updated_a as (update appointments set doctor_id=$2, clinic_id=$3, start_time=$4, end_time=$5 where id=$1 returning *)

	select updated_a.id, updated_a.doctor_id, updated_a.clinic_id, updated_a.start_time, updated_a.end_time, 
	doctors.id, doctors.name, doctors.specialty, clinics.id, clinics.cname, clinics.location
	from updated_a
	join doctors on updated_a.doctor_id = doctors.id
	join clinics on updated_a.clinic_id = clinics.id`

	row := m.DB.QueryRow(query, id, &a.DoctorID, &a.ClinicID, &a.StartTime, &a.EndTime)
	
	err := row.Scan(&a.ID, &a.DoctorID, &a.ClinicID, &a.StartTime, &a.EndTime, 
		&a.Doctor.ID, &a.Doctor.Name, &a.Doctor.Specialty,
		&a.Clinic.ID, &a.Clinic.Name, &a.Clinic.Location)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &a, nil
}

func (m *postgresDBRepo) DeleteAppointment(id int) (*sql.Result, error) {
	query := `delete from appointments where id=$1`

	res, err := m.DB.Exec(query, id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &res, nil
}
