package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/B0go/fin/database"
	"github.com/B0go/fin/env"
	"github.com/apex/log"
	"github.com/gorilla/mux"
)

func main() {
	cfg, err := env.Get()
	if err != nil {
		log.WithError(err).
			Fatal("failed to load config")
	}

	_, err = database.Connect(cfg)
	if err != nil {
		log.WithError(err).
			Fatal("failed to connect to db")
	}

	log.Info("starting fin")
	mux := mux.NewRouter()

	mux.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "OK") // nolint: gas
	})

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.WithError(err).Error("server startup failed")
	}
}
