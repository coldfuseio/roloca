package main

import (
	"log"
	"net/http"
	"time"
  "github.com/jpillora/go-ogle-analytics"
)

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		
		client, err := ga.NewClient("UA-67119922-8")
    if err != nil {
      panic(err)
    }
    err = client.Send(ga.NewEvent(name, r.RequestURI))
    if err != nil {
        panic(err)
    }

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
