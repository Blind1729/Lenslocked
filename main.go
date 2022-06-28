package main

import (
	"Lenslocked/views"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var (
	homeView    *views.View
	contactView *views.View
)

func home(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeView.Render(w, nil); err != nil {
		panic(err)
	}
}

func faq(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	_, err := fmt.Fprintf(w, "<h1>Frequently asked questions</h1>")
	if err != nil {
		return
	}
}

func pageDoesNotExist(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	_, err := fmt.Fprintf(w, "<h2>Sorry, the page does not exist</h2>")
	if err != nil {
		return
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
	router := mux.NewRouter()
	router.HandleFunc("/", home)
	router.HandleFunc("/faq", faq)
	router.HandleFunc("/contact", contact)
	router.NotFoundHandler = http.HandlerFunc(pageDoesNotExist)
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		return
	}
}
