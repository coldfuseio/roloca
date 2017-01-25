package main

// import (
//   "fmt"
// )

type County struct {
	Auto      string    `json:"auto"`
	Name      string    `json:"nume"`
	Region    string    `json:"regiune"`
}

type CountySimple struct {
	Auto      string    `json:"auto"`
	Name      string    `json:"nume"`
}

func GetAllCounties() (counties []CountySimple) {
  db := Connect()
  defer db.Close()

  rows, err := db.Query("select auto, nume from counties")
	checkErr(err)

	defer rows.Close()
	for rows.Next() {
		var county CountySimple
		rows.Scan(&county.Auto,	&county.Name)
		counties = append(counties, county)
	}
	return counties
}
