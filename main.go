package main

import (
    "log"
    "net/http"
	"strings"
    "github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(`{"message": "welcome to go"}`))
    
}
func hello(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    w.Write([]byte(`{"message": "Hello world!"}`))
}


func customhello(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	vars := mux.Vars(r)
	name := strings.Title(strings.ToLower(vars["name"]))
	
    w.Write([]byte(`{"message": "Hello ` + name  +`!"}`))
}


func main() {
    r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/hello", hello)
	r.HandleFunc("/greet/{name}", customhello)
    log.Fatal(http.ListenAndServe(":8080", r))
}