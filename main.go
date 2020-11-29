package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	Articles = []Article{
		Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	handleRequests()
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnAllArticles)
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)

	//http.HandleFunc("/", homePage)
	//http.HandleFunc("/articles", returnAllArticles)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func returnSingleArticle(writer http.ResponseWriter, request *http.Request) {
	//defer recoveryFunction(writer)
	vars := mux.Vars(request)
	key := vars["id"]
	found := false
	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(writer).Encode(article)
			found = true
		}
	}
	if found == false {
		writer.WriteHeader(404)
	}

}

func recoveryFunction(writer http.ResponseWriter) {
	fmt.Fprintf(writer, "panic triggered")
}

func homePage(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "welcome to the homepage with method: ", request.Method)
	fmt.Println("endpoint hit: homepage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article
