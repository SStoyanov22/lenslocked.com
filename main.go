package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Someone visited our page", r.URL)
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, "<h1>Welcome to my super awsome site!</h1>")
}
func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe("localhost:3000", nil)
}
