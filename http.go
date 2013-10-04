package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"mhacks2013f/bass"
)

func handler(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	country_name := vars["country_name"]
	c := bass.NewRDB()
	str := bass.PullRDB(country_name, c)
	for _, s := range str {
		fmt.Fprintf(rw, "%s", s)
	}
	bass.CloseRDB(c)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/country/{country_name}", handler)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
