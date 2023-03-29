package dbrepo

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/Wai30Yan/cna-server/model"
)

func (m *postgresDBRepo) GetAllSchedules() []model.Schedule {
	query := "select id, doctor_name, start_time, end_time from schedules"
	rows, err := m.DB.Query(query)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer rows.Close()

	var schedules []model.Schedule

	for rows.Next() {
		var s model.Schedule
		var start, end string
		err := rows.Scan(&s.ID, &s.DoctorName, &start, &end)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		st, err := time.Parse(time.RFC3339, start)
		s.StartTime = st.Format("02-01-06 15:04")
		if err != nil {
			log.Fatal(err)
			return nil
		}
		et, err := time.Parse(time.RFC3339, end)
		s.EndTime = et.Format("02-01-06 15:04")
		if err != nil {
			log.Fatal(err)
			return nil
		}
		schedules = append(schedules, s)
		fmt.Println(s)
	}

	return schedules
}

func (m *postgresDBRepo) GetScheduleByID(id int) *model.Schedule {
	var s model.Schedule
	query := `select * from schedules where id=$1`

	var start, end string

	err := m.DB.QueryRow(query, id).Scan(&s.ID, &s.DoctorName, &start, &end)
    if err != nil {
        log.Fatal(err)
    }

	st, err := time.Parse(time.RFC3339, start)
	s.StartTime = st.Format("02-01-06 15:04")
	if err != nil {
		log.Fatal(err)
		return nil
	}

	et, err := time.Parse(time.RFC3339, end)
	s.EndTime = et.Format("02-01-06 15:04")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return &s
}

func (m *postgresDBRepo) InsertSchedule(s *model.Schedule) *model.Schedule {
	query := `insert into schedules (doctor_name, start_time, end_time)
		values ($1, $2, $3)
		returning id, doctor_name, start_time, end_time
		`

	rows, err := m.DB.Query(query, s.DoctorName, s.StartTime, s.EndTime)
	if err != nil {
		log.Print(err)
		return nil
	}

	var newS model.Schedule
	if rows.Next() {
		var start, end string
		err := rows.Scan(&newS.ID, &newS.DoctorName, &start, &end)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		st, err := time.Parse(time.RFC3339, start)
		newS.StartTime = st.Format("02-01-06 15:04")
		if err != nil {
			log.Fatal(err)
			return nil
		}
		et, err := time.Parse(time.RFC3339, end)
		newS.EndTime = et.Format("02-01-06 15:04")
		if err != nil {
			log.Fatal(err)
			return nil
		}		
	}

	return &newS
}

func (m *postgresDBRepo) UpdateSchedule(id int, updated *model.Schedule) *model.Schedule {
	var newS model.Schedule

	query := `update schedules set doctor_name=$1, start_time=$2, end_time=$3 where id=$4
		returning id, doctor_name, start_time, end_time
	`

	rows, err := m.DB.Query(query, updated.DoctorName, updated.StartTime, updated.EndTime, id)
	if err != nil {
		log.Print(err)
		return nil
	}

	if rows.Next() {
		var start, end string
		err := rows.Scan(&newS.ID, &newS.DoctorName, &start, &end)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		st, err := time.Parse(time.RFC3339, start)
		newS.StartTime = st.Format("02-01-06 15:04")
		if err != nil {
			log.Fatal(err)
			return nil
		}
		et, err := time.Parse(time.RFC3339, end)
		newS.EndTime = et.Format("02-01-06 15:04")
		if err != nil {
			log.Fatal(err)
			return nil
		}	
	}
	
	return &newS
}

func (m *postgresDBRepo) DeleteSchedule(id int) sql.Result {
	query := `delete from schedules where id=$1`

	res, err := m.DB.Exec(query, id)
	if err != nil {
		log.Print(err)
	}

	return res
}
