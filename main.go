package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	handleRequests()
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func homePage(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "welcome to the homepage with method: ", request.Method)
	fmt.Println("endpoint hit: homepage")
}

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article
