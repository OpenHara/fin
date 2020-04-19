package database

import (
	"fmt"

	"github.com/B0go/fin/env"
	"github.com/B0go/fin/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func Connect(cfg *env.Config) (*sqlx.DB, error) {
	return sqlx.Open("mysql", fmt.Sprintf("root:%s@(%s)/fin", cfg.DatabasePassword, cfg.DatabaseUrl))
}

func Persist(db *sqlx.DB, e *model.Entry) error {
	tx := db.MustBegin()
	_, err := tx.NamedExec("INSERT INTO entries (id, description, value_in_cents, type) VALUES (:id, :description, :value_in_cents, :type)", e)
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}
