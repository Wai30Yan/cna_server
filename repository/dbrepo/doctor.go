package dbrepo

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
	"github.com/Wai30Yan/cna-server/model"
)

func (m *postgresDBRepo) GetAllDoctors() ([]*model.Doctor, error) {
	var doctors []*model.Doctor 
	query := `select * from doctors`
	rows, err := m.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var doc model.Doctor
		err := rows.Scan(&doc.ID, &doc.Name, &doc.Specialty)
		if err != nil {
			return nil, err
		}

		doctors = append(doctors, &doc)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	
	return doctors, nil
}


func (m *postgresDBRepo) GetDoctorByID(id int) (*model.Doctor, error) {
	var doctor model.Doctor

	query := `select * from doctors where id=$1`

	row, err := m.DB.Query(query, id)
	if err != nil {
		fmt.Println(err)
        return nil, fmt.Errorf("error inserting doctor: %v", err)
    }

	if row.Next() {
		err := row.Scan(&doctor.ID, &doctor.Name, &doctor.Specialty)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}

	return &doctor, nil
}

func (m *postgresDBRepo) InsertDoctor(doctor *model.Doctor) (*model.Doctor, error) {
    query := `INSERT INTO Doctors(name, specialty) VALUES ($1, $2) RETURNING id, name, specialty`

    row, err := m.DB.Query(query, doctor.Name, doctor.Specialty)
    if err != nil {
		fmt.Println(err)
        pgErr, ok := err.(*pq.Error)
        if ok && pgErr.Code.Name() == "unique_violation" {
            return nil, fmt.Errorf("doctor with name %s already exists", doctor.Name)
        }
        return nil, fmt.Errorf("error inserting doctor: %v", err)
    }
	var newDoc model.Doctor
	if row.Next() {
		err := row.Scan(&newDoc.ID, &newDoc.Name, &newDoc.Specialty)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}

    return &newDoc, nil
}

func (m *postgresDBRepo) UpdateDoctor(id int, doctor *model.Doctor) (*model.Doctor, error) {
	query := `update doctors set name=$2, specialty=$3 where id=$1`

	row := m.DB.QueryRow(query, id, &doctor.Name, &doctor.Specialty)
	err := row.Scan(&doctor.ID, &doctor.Name, &doctor.Specialty)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return doctor, nil
}

func (m *postgresDBRepo) DeleteDoctor(id int) (*sql.Result, error) {
	query := `delete from doctors where id=$1`
	res, err := m.DB.Exec(query, id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &res, nil
}


// select * from doctors where name ilike'%th%'

// SELECT *
// FROM appointments
// JOIN doctors ON appointments.doctor_id = doctors.id
// WHERE doctors.name ILIKE '%th%'

// SELECT *
// FROM appointments
// JOIN doctors ON appointments.doctor_id = doctors.id
// WHERE doctors.id = 4

// select * from appointments a join doctors d on d.id=a.doctor_id where d.id=4