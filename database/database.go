package database

import (
	"fmt"

	"github.com/B0go/fin/env"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func Connect(cfg *env.Config) (*sqlx.DB, error) {
	return sqlx.Open("mysql", fmt.Sprintf("root:%s@(%s)/fin", cfg.DatabasePassword, cfg.DatabaseUrl))
}
