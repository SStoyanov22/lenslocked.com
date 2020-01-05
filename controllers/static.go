package controllers

import "github.com/SStoyanov22/lenslocked.com/views"

type Static struct {
	HomeView *views.View
	ContactView *views.View
}

func NewStatic() *Static {
	return &Static{
		HomeView: views.NewView("bootstrap","views/static/home.gohtml"),
		ContactView: views.NewView("bootstrap","views/static/contact.gohtml"),
	}
}