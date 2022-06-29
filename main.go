package main

import (
	"Lenslocked/controllers"
	"Lenslocked/views"
	"github.com/gorilla/mux"
	"net/http"
)

var (
	homeView     *views.View
	contactView  *views.View
	faqView      *views.View
	pageNotFound *views.View
)

func home(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeView.Render(w, nil); err != nil {
		panic(err)
	}
}

func faq(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := faqView.Render(w, nil); err != nil {
		panic(err)
	}
}

func pageDoesNotExist(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := pageNotFound.Render(w, nil); err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := contactView.Render(w, nil); err != nil {
		panic(err)
	}
}

func main() {
	homeView = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")
	faqView = views.NewView("bootstrap", "views/faq.gohtml")
	usersC := controllers.NewUsers()
	pageNotFound = views.NewView("bootstrap", "views/resource-not-found.gohtml")
	router := mux.NewRouter()
	router.HandleFunc("/", home)
	router.HandleFunc("/faq", faq)
	router.HandleFunc("/contact", contact)
	router.HandleFunc("/signup", usersC.New).Methods("GET")
	router.HandleFunc("/signup", usersC.Create).Methods("POST")
	router.NotFoundHandler = http.HandlerFunc(pageDoesNotExist)
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		return
	}
}
