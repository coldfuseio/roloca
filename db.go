package main

import (
	"os"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"fmt"
)

func Connect() (db *sql.DB) {
	if _, err := os.Stat("/loc.db"); os.IsNotExist(err) {
  	fmt.Println("existaaaa")
	}
	db, err := sql.Open("sqlite3", "loc.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
