package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/robfig/cron"
)

func start(gateway string) {
	routes := mux.NewRouter()
	//Routes
	routes.HandleFunc("/update", updateHandler).
		Methods("POST")

	log.Info("Server running on ", gateway)
	log.Panic(http.ListenAndServe(gateway, routes))
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		w.WriteHeader(http.StatusInternalServerError)
	}()

	crono.Stop()

	log.Info("Configuration file config.yaml updated successfully.")

	insta.Logout()
	go forever()
	w.WriteHeader(http.StatusOK)
}

func schedule() {
	crono = cron.New()

	crono.AddFunc(
		yamlOpt.Daemon.Interval,
		func() {
			getTagIds()
		},
	)

	crono.Start()
}
