package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title   string "json:\"Title\""
	Desc    string "json:\"Desc\""
	Content string "json:\"Content\""
}
type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Test Title", Desc: "Test Desc", Content: "Hello World"},
	}
	fmt.Println("EndPoint Hit: All Articles EndPoint")
	json.NewEncoder(w).Encode(articles)
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage EndPoint Hit")
}
func handleRequests() {
	http.HandleFunc("/", homepage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

/* func NewReq() req {
	temp := new(req)
	temp.datas = make(map[string]string)
	temp.datas["0"] = "zero"
	temp.datas["1"] = "one"

	return *temp
}
func (r req) Get(key string) {

}
func (r req) Set(key string, value string) {

}
func (r req) Flush() {

} */
