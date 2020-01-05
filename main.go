package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/SStoyanov22/lenslocked.com/controllers"
)

func main() {
	staticC := controllers.NewStatic()	
	usersC := controllers.NewUsers()
	
	r := mux.NewRouter()
	r.Handle("/",staticC.HomeView).Methods("GET")
	r.Handle("/contact",staticC.ContactView).Methods("GET")
	r.HandleFunc("/signup",usersC.New).Methods("GET")
	r.HandleFunc("/signup",usersC.Create).Methods("GET", "POST")
	http.ListenAndServe(":3000", r)
}

func must (err error){
	if err != nil {
		panic(err)
	}
}