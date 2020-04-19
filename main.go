package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/B0go/fin/env"
	"github.com/apex/log"
	"github.com/gorilla/mux"
)

func main() {
	_, err := env.MustGet()
	if err != nil {
		log.WithError(err).
			Fatal("failed to load config")
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
