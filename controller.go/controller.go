package controller

import (
	"encoding/json"
	"net/http"

	"github.com/B0go/fin/database"
	"github.com/B0go/fin/model"
	"github.com/apex/log"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func CreateEntry(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var e model.Entry
		if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
			log.WithError(err).
				Error("decoding request json into struct")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		e.Id = uuid.New().String()

		err := database.Persist(db, &e)
		if err != nil {
			log.WithError(err).
				Error("failed to persist")
		}

		err = json.NewEncoder(w).Encode(e)
		if err != nil {
			log.WithError(err).
				Error("encoding struct into json")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
