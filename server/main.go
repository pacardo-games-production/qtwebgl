package main

import (
	"log"
	"net/http"
	"text/template"
)

func home(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("index.html")
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Error", 500)
		return
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	fileServer := http.FileServer(http.Dir("./Build/"))
	mux.Handle("/Build/", http.StripPrefix("/Build", fileServer))

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
