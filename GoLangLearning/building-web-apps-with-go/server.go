package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/russross/blackfriday"
    "io/ioutil"
)

func main(){

	r := mux.NewRouter().StrictSlash(false)
	// Posts Collection
	posts := r.Path("/posts").Subrouter()
	posts.Methods("GET").HandlerFunc(PostsIndexHandler)
	posts.Methods("POST").HandlerFunc(PostsCreateHandler)

	// Posts singular
	post := r.PathPrefix("/posts/edit/{id}").Subrouter()
	post.Methods("GET").Path("/edit").HandlerFunc(PostEditHandler)
	post.Methods("GET").HandlerFunc(PostShowHandler)
	post.Methods("POST", "PUT").HandlerFunc(PostUpdateHandler)
	post.Methods("DELETE").HandlerFunc(PostDeleteHandler)

	// markdown
	markdown := r.Path("/markdown").Subrouter()
	markdown.Methods("GET").HandlerFunc(GetMarkdown)
	markdown.Methods("POST").HandlerFunc(GenerateMarkdown)		
	r.HandleFunc("/", HomeHandler)

	http.ListenAndServe(":8080", r)
}

func HomeHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Home"))
}
func PostsIndexHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("posts index"))
}
func PostsCreateHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("posts create"))
}
func PostShowHandler(rw http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	rw.Write([]byte("showing post " + id))
}
func PostUpdateHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("post update"))
}
func PostDeleteHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("post delete"))
}
func PostEditHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("post edit"))
}


func GetMarkdown(w http.ResponseWriter, r *http.Request){
	data, err := ioutil.ReadFile("./public/index.html")

	if err != nil{
		w.Write([]byte(err.Error()))
	}
	fmt.Println(string(data))
	w.Write(data)		
}

func GenerateMarkdown(w http.ResponseWriter, r *http.Request){
	markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
	w.Write(markdown)		
}
