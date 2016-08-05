package main

import (
	"encoding/json"
	"net/http"
)


func CountiesAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
  err := json.NewEncoder(w).Encode(GetAllCounties())
  checkErr(err)
}
