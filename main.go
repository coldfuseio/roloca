package main // import "github.com/nalexeric/roloca"

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Roloca! \n /judete - list all counties in Romania \n /orase - list all cities in Romania \n /orase/<county-code> - list all cities in a county (e.g. /orase/GL - all cities in Galati county)"))
}
