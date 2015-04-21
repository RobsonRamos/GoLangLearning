package main

import (
	"flag"
	"log"
	"net/http"
	"text/template"	
	"path/filepath"
	"sync"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/facebook"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/google"
)

type templateHandler struct {
	once 		sync.Once
	filename 	string
	templ		*template.Template
}


func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	t.once.Do(func(){
		t.templ = template.Must(template.ParseFiles( filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, r)
}

func main(){
	var addr = flag.String("addr", ":8080", "The addr of the application")
	flag.Parse()

	gomniauth.SetSecurityKey("yLiCQYG7CAflDavqGH461IO0MHp7TEbpg6TwHBWdJzNwYod1i5ZTbrIF5bEoO3oP") // NOTE: DO NOT COPY THIS - MAKE YOR OWN!
	gomniauth.WithProviders(
		github.New("3d1e6ba69036e0624b61", "7e8938928d802e7582908a5eadaaaf22d64babf1", "http://localhost:8080/auth/github/callback"),
		google.New("1051709296778.apps.googleusercontent.com", "7oZxBGwpCI3UgFMgCq80Kx94", "http://localhost:8080/auth/google/callback"),
		facebook.New("537611606322077", "f9f4d77b3d3f4f5775369f5c9f88f65e", "http://localhost:8080/auth/facebook/callback"),
	)

	r := NewRoom()

	//http.Handle("/", &templateHandler{ filename : "chat.html" })
	http.Handle("/chat", MustAuth(&templateHandler{ filename : "chat.html"}))
	http.Handle("/login", &templateHandler{ filename : "login.html"})	
	http.HandleFunc("/auth", loginHandler)
	http.Handle("/room", r)

	go r.run()
	log.Println("Starting web server on", *addr)

	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}	
}
