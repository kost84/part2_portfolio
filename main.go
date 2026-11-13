package main

import (
	//"fmt"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main(){

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", home)

	http.HandleFunc("/acceuil", acceuil)	

	http.HandleFunc("/project", project)

	fmt.Println("serveur lancer sur http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil{
		log.Fatal(err)
	}
}

func home(w http.ResponseWriter, r *http.Request){
	renderTemplate(w, "page/index.html")
}

func acceuil(w http.ResponseWriter, r *http.Request){
	renderTemplate(w, "page/acceuil.html")
}

func project(w http.ResponseWriter, r *http.Request){
	renderTemplate(w, "page/project.html")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles(tmpl)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}