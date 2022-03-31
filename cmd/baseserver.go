package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ok", returnOk)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func returnOk(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnOk")
	json.NewEncoder(w).Encode("ok")
}

func main() {
	handleRequests()
}
