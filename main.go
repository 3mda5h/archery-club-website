package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.ParseGlob("pages/*.html"))

func renderTemplate(w http.ResponseWriter, page string) {
	err := templates.ExecuteTemplate(w, page, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// Route handlers
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "index.html")
	})
	http.HandleFunc("/practices", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "practices.html")
	})
	http.HandleFunc("/schedule", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "schedule.html")
	})
	http.HandleFunc("/membership", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "membership.html")
	})
	http.HandleFunc("/locations", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "locations.html")
	})
	http.HandleFunc("/photo_gallery", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "photo_gallery.html")
	})
	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "contact.html")
	})

	log.Println("Server running at http://localhost:5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
