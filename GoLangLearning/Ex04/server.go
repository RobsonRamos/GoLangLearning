package main

import(
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request){
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusOK)
			fmt.Println(w, "Hello!")
		},
	)
	http.ListenAndServe(":8080", nil)
}
