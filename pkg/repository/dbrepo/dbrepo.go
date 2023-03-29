package dbrepo

import (
	"database/sql"

	"github.com/Wai30Yan/cna-server/pkg/repository"
	"github.com/Wai30Yan/cna-server/pkg/config"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB: conn,
	}
}