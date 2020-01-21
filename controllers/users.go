package controllers

import (
	"fmt"
	"github.com/SStoyanov22/lenslocked.com/models"
	"github.com/SStoyanov22/lenslocked.com/views"
	"net/http"
)

//used to create new users controller
//will panic if templates are not passed correctly
//and should only be used during initial setup
func NewUsers(us *models.UserService) *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "users/new"),
		us:      us,
	}
}

type Users struct {
	NewView *views.View
	us      *models.UserService
}

func New(u *Users, w http.ResponseWriter, r *http.Request) {

}

type SignupForm struct {
	Name     string `shhema:"name"`
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

//GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

//Create is used to process the signup form when a user submits it.
//This is used to create a new users account
//
//POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var form SignupForm
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}

	user := models.User{
		Name:  form.Name,
		Email: form.Email,
	}

	if err := u.us.Create(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "User is", user)
}
