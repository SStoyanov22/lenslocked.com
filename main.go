package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func handlerFunc(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my awsome site!</h1>")
}

func homeHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>To get in touch, please send and email to <a href = \"support@lenslocked.com\"</h1>")
}

func faqHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>FAQ Page</h1>")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>Sorry, but we couldn't find the page you were looking for</h1>")
}

func main() {
	r := httprouter.New()
	r.GET("/", homeHandler)
	r.GET("/contact", contactHandler)
	r.GET("/faq", faqHandler)
	r.NotFound = http.HandlerFunc(notFoundHandler)
	http.ListenAndServe(":8000", r)
}
