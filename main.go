package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func home(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	_, err := fmt.Fprintf(w, "<h1>Welcome to my site</h1>")
	if err != nil {
		return
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

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", home)
	router.HandleFunc("/faq", faq)
	router.NotFoundHandler = http.HandlerFunc(pageDoesNotExist)
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		return
	}
}
