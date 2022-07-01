// static controller renders all the static pages

package controllers

import "Lenslocked/views"

func NewStatic() *Static {
	return &Static{
		HomeView:     views.NewView("bootstrap", "static/home"),
		ContactView:  views.NewView("bootstrap", "static/contact"),
		FaqView:      views.NewView("bootstrap", "static/faq"),
		PageNotFound: views.NewView("bootstrap", "static/resource-not-found"),
	}
}

type Static struct {
	HomeView     *views.View
	ContactView  *views.View
	FaqView      *views.View
	PageNotFound *views.View
}
