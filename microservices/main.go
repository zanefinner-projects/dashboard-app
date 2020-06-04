package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zanefinner-projects/dashboard-app/accounts"
)

type message struct {
	Content string
}

func main() {
	fmt.Println("start")
	defer fmt.Println("finish")

	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		defer fmt.Println("GET: [/]")
	}).Methods("GET")

	//##ACCOUNTS
	r.HandleFunc("/accounts/login", accounts.Login).Methods("POST")

	//##OTHER
	r.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		defer fmt.Println("POST: [/echo]")
		var m message
		err := json.NewDecoder(r.Body).Decode(&m)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		go fmt.Println("Retrieved", m.Content)
		go fmt.Println("Sent Back")
		w.Write([]byte(m.Content))

	}).Methods("POST")
	log.Fatal(http.ListenAndServe(":8001", r))
}
