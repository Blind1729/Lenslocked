// static controller renders all the static pages

package controllers

import "Lenslocked/views"

func NewStatic() *Static {
	return &Static{
		HomeView:     views.NewView("bootstrap", "views/static/home.gohtml"),
		ContactView:  views.NewView("bootstrap", "views/static/contact.gohtml"),
		FaqView:      views.NewView("bootstrap", "views/static/faq.gohtml"),
		PageNotFound: views.NewView("bootstrap", "views/static/resource-not-found.gohtml"),
	}
}

type Static struct {
	HomeView     *views.View
	ContactView  *views.View
	FaqView      *views.View
	PageNotFound *views.View
}
