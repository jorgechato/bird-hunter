package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func start(gateway string) {
	routes := mux.NewRouter()
	//Routes
	routes.HandleFunc("/update", updateHandler).
		Methods("POST")

	log.Fatal(http.ListenAndServe(gateway, routes))
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		w.WriteHeader(http.StatusInternalServerError)
	}()

	readYaml(*configFilePtr, &yamlOpt)
	w.WriteHeader(http.StatusOK)
}
