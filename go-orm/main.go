package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HellowWorld API is accesed")
	fmt.Fprintf(w, "HellowWorld API is accesed")
}

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", helloWorld).Methods("GET")
	myRouter.HandleFunc("/users", AllUsers)
	myRouter.HandleFunc("/user/{name}/{email}", NewUsers).Methods("POST")
	myRouter.HandleFunc("/user/{name}", DeleteUsers).Methods("DELETE")
	myRouter.HandleFunc("/user/{name}/{email}", UpdateUsers).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	fmt.Println("Go ORM ")
	handleRequest()
}
