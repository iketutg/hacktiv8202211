package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

type Data struct {
	Title string
	Body  string
}

func renderTemplates(w http.ResponseWriter, data *Data, navigasi string) {
	err := templates.ExecuteTemplate(w, navigasi, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	body := "This is index Golang Pages"
	title := "Index Page"
	data := &Data{
		Title: title,
		Body:  body,
	}
	renderTemplates(w, data, "index")
}
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	body := "This is About Golang Pages"
	title := "About Page"
	data := &Data{
		Title: title,
		Body:  body,
	}
	renderTemplates(w, data, "index")

}
func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)
	log.Println("Server Up : 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
