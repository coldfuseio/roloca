package main // import "github.com/nalexeric/roloca"

import (
	"log"
	"net/http"
	"fmt"
	"github.com/yosssi/gohtml"
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
	h := `
	<!DOCTYPE html><html>
	<head>
		<title>Roloca</title>
	</head>
	<body>
		Roloca! <br>
		<a href='/judete'>/judete</a> - list all counties in Romania <br>
		<a href='/orase'>/orase</a> - list all cities in Romania <br>
		/orase/:countyCode - list all cities in a county (e.g. <a href='/orase/GL'>/orase/GL</a> - all cities in Galati county)
		<br><br>
		developed by <a href='http://coldfuse.io'>coldfuse.io</a>,<br>using data from <a href='https://github.com/romania/localitati'>github.com/romania/localitati</a>
	</body>
	`
	fmt.Fprintf(w, gohtml.Format(h))
}
