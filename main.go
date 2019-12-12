package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Header", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcome to my super awsome site!</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "<a href=\"support@lenslocked.com\">support@lenslocked.com</a/")
	}

}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe("localhost:8000", nil)
}
