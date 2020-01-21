package main

import (
	"fmt"
	"github.com/SStoyanov22/lenslocked.com/controllers"
	"github.com/SStoyanov22/lenslocked.com/models"
	"github.com/gorilla/mux"
	"net/http"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "your-password"
	dbname   = "lenslocked_dev"
)

func main() {
	// Create a DB connection string and then use it to
	// create our model services.
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	us, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer us.Close()
	us.Automigrate()

	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers(us)

	r := mux.NewRouter()
	r.Handle("/", staticC.HomeView).Methods("GET")
	r.Handle("/contact", staticC.ContactView).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("GET", "POST")
	http.ListenAndServe(":3000", r)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
