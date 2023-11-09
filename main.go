package main

import (
	api "github.com/spaceox/stalewall/api"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		api.Handler(w, r)
	})
	println("Serving on http://localhost:3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}
