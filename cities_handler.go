package main

import (
	"encoding/json"
	"net/http"
  "github.com/gorilla/mux"
	"strings"
)


func CitiesAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(GetAllCities())
  checkErr(err)
}


func CitiesInCounty(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	county := vars["county"]

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(GetCitiesInCounty(strings.ToUpper(county)))
	checkErr(err)
}
