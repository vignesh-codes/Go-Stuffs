package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}
type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {

	articles := Articles{
		Article{Title: "Title", Desc: "Descc", Content: "Hellow"},
	}

	fmt.Println("Endpoint Hit: All Articles")
	json.NewEncoder(w).Encode(articles)
}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}
func allArti(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the allArti!")
	fmt.Println("Endpoint Hit: allArti")
}

func handleRequests() {
	//using this to only be accessable with post or get req
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", allArti).Methods("POST")

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func mains() {
	handleRequests()
}
