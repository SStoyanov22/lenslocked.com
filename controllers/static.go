package controllers

import "github.com/SStoyanov22/lenslocked.com/views"

type Static struct {
	HomeView *views.View
	ContactView *views.View
}

func NewStatic() *Static {
	return &Static{
		HomeView: views.NewView("bootstrap","static/home"),
		ContactView: views.NewView("bootstrap","static/contact"),
	}
}