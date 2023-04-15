package dbrepo

import (
	"fmt"

	"github.com/Wai30Yan/cna-server/pkg/model"
)

func (m *postgresDBRepo) SignUp(username, password string) (*model.Admin, error) {
	query := `insert into admins (username, password) values ($1,$2) returning *`

	var admin *model.Admin

	row := m.DB.QueryRow(query, username, password)
	err := row.Scan(&admin.ID, &admin.UserName, &admin.Password)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return admin, nil
}

func (m *postgresDBRepo) LogIn(username, password string) (*model.Admin, error) {
	return nil, nil
}

func (m *postgresDBRepo) LogOut() {

}
