package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	//r.HandleFunc("", HomeHandler)
	//r.HandleFunc("", PerfHandler)
	//r.HandleFunc("", AuthHandler)
	http.Handle("/", r)
}
