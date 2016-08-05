package main

// import (
//   "fmt"
// )

type City struct {
	Id        	int       `json:"id"`
	County    	string    `json:"judet"`
	Name      	string    `json:"nume"`
	NameSimple	string    `json:"simplu,omitempty"`
	Admin     	string    `json:"comuna,omitempty"`
	Long      	float64   `json:"long"`
	Lat       	float64   `json:"lat"`
  Siruta    	string    `json:"siruta"`
}

type CitySimple struct {
  Name      	string    `json:"nume"`
  NameSimple  string    `json:"simplu,omitempty"`
  Admin     	string    `json:"comuna,omitempty"`
}

func GetAllCities() (cities []CitySimple) {
  db := Connect()
  defer db.Close()

  rows, err := db.Query("select nume, numeSimplu, comuna from cities")
	checkErr(err)

	defer rows.Close()
	for rows.Next() {
		var city CitySimple
		rows.Scan(
      &city.Name,
			&city.NameSimple,
      &city.Admin,
    )
		cities = append(cities, city)
	}
	return cities
}

func GetCitiesInCounty(county string) (cities []CitySimple) {
  db := Connect()
  defer db.Close()

  stmt, err := db.Prepare("select nume, numeSimplu, comuna from cities where judet=? order by comuna asc, nume asc")
  checkErr(err)

  rows, err := stmt.Query(county)
  checkErr(err)

  defer rows.Close()
  for rows.Next() {
		var city CitySimple
		rows.Scan(
      &city.Name,
			&city.NameSimple,
      &city.Admin,
    )
		cities = append(cities, city)
	}
	return cities
}
