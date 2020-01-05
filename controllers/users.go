package controllers

import (
"github.com/SStoyanov22/lenslocked.com/views"
"net/http"
"fmt"
)

//used to create new users controller
//will panic if templates are not passed correctly
//and should only be used during initial setup
func NewUsers() *Users{
	return &Users{
		NewView: views.NewView("bootstrap","users/new"),
	}
}

type Users struct{
	NewView *views.View
}

func New(u *Users, w http.ResponseWriter, r *http.Request){

}

type SignupForm struct {
	Email string  `schema:"email"`
	Password string `schema:"password"`
}

//GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request){
	if err := u.NewView.Render(w,nil); err!= nil {
		panic(err)
	}
}

//Create is used to process the signup form when a user submits it.
//This is used to create a new users account
//
//POST /signup
func (u *Users) Create (w http.ResponseWriter, r *http.Request){
	var form SignupForm
	if err := parseForm(r,&form); err!=nil{
		panic(err)
	}
	fmt.Fprintln(w, form)
}