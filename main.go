package main

import (
	"Lenslocked/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers()
	router := mux.NewRouter()
	router.Handle("/", staticC.HomeView).Methods("GET")
	router.Handle("/faq", staticC.FaqView).Methods("GET")
	router.Handle("/contact", staticC.ContactView).Methods("GET")
	router.HandleFunc("/signup", usersC.New).Methods("GET")
	router.HandleFunc("/signup", usersC.Create).Methods("POST")
	router.NotFoundHandler = http.Handler(staticC.PageNotFound)
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		return
	}
}
