package dbrepo

import (
	"database/sql"
	"fmt"

	"github.com/Wai30Yan/cna-server/pkg/model"
)

func (m *postgresDBRepo) GetAllClinics() ([]*model.Clinic, error) {
	var clinics []*model.Clinic

	query := `select * from clinics`
	row, err := m.DB.Query(query)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for row.Next() {
		var c model.Clinic
		err = row.Scan(&c.ID, &c.Name, &c.Location)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		clinics = append(clinics, &c)
	}
	return clinics, nil
}

func (m *postgresDBRepo) GetClinicByID(id int) (*model.Clinic, error) {
	var clinic model.Clinic

	query := `select * from clinics where id=$1`

	row, err := m.DB.Query(query, id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if row.Next() {
		err := row.Scan(&clinic.ID, &clinic.Name, &clinic.Location)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}

	return &clinic, nil
}

func (m *postgresDBRepo) InsertClinic(clinic *model.Clinic) (*model.Clinic, error) {
	query := `insert into clinics (cname, location) values ($1,$2) returning *`

	row := m.DB.QueryRow(query, &clinic.Name, &clinic.Location)
	err := row.Scan(&clinic.ID, &clinic.Name, &clinic.Location)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return clinic, nil
}
func (m *postgresDBRepo) UpdateClinic(id int, clinic model.Clinic) (*model.Clinic, error) {
	query := `update clinics set cname=$2, location=$3 where id=$1`

	row := m.DB.QueryRow(query, id, &clinic.Name, &clinic.Location)

	err := row.Scan(&clinic.ID, &clinic.Name, &clinic.Location)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &clinic, nil
}
func (m *postgresDBRepo) DeleteClinic(id int) (*sql.Result, error) {
	query := `delete from clinics where id=$1`

	res, err := m.DB.Exec(query, id)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &res, nil
}