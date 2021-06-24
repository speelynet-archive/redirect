package main

import (
	"net/http"
	unit "unit.nginx.org/go"
)

func main() {
	if e := unit.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, Redirect(r.URL), http.StatusMovedPermanently)
	})); e != nil {
		panic(e)
	}
}
