package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Home",
		"Get",
		"/",
		HomeHandler,
	},
	Route{
		"HomeRo",
		"Get",
		"/ro",
		HomeRoHandler,
	},

	Route{
		"CountiesAll",
		"GET",
		"/judete",
		CountiesAll,
	},
	Route{
		"CitiesAll",
		"GET",
		"/orase",
		CitiesAll,
	},
	Route{
		"CitiesInCounty",
		"GET",
		"/orase/{county}",
		CitiesInCounty,
	},
}
