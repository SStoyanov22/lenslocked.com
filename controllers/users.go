package controllers

import (
"github.com/SStoyanov22/lenslocked.com/views"
"github.com/gorilla/schema"
"net/http"
"fmt"
)

//used to create new users controller
//will panic if templates are not passed correctly
//and should only be used during initial setup
func NewUsers() *Users{
	return &Users{
		NewView: views.NewView("bootstrap","views/users/new.gohtml"),
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
	if err := r.ParseForm()	; err != nil {
		panic(err)
	}

	dec := schema.NewDecoder()
	var form SignupForm
	if err:= dec.Decode(&form, r.PostForm); err!=nil{
		panic(err)
	}

	fmt.Fprintln(w, form)
	fmt.Fprintln(w, r.PostForm["email"])
	fmt.Fprintln(w, r.PostForm["password"])
	fmt.Fprintln(w, r.PostFormValue("email"))
	fmt.Fprintln(w,"This is temporary response.")
}