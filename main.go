package main

import (
	"fmt"
	"net/http"
	"log"
    "github.com/julienschmidt/httprouter"
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

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Fprintf(w, "Hola, %s!\n", ps.ByName("name"))
}

func main() {
	router := httprouter.New()
	router.GET("/hello/:name/spanish", Hello)

    log.Fatal(http.ListenAndServe(":8000", router))
}
