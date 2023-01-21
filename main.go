package main

import (
    "fmt"
    api "github.com/spaceox/stalewall/api"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		api.Handler(w, r)
	})
	err := http.ListenAndServe(":3000", nil)
    fmt.Println("Serving on http://localhost:3000")
    if err != nil {
        panic(err)
    }
}
