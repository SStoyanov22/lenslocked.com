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
		fmt.Fprint(w, "<a href=\"support@lenslocked.com\">support@lenslocked.com</a>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>Page not found</h1><p>Please contact us if you keep being sent to an invalid page</p>")
	}

}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe("localhost:8000", nil)
}
