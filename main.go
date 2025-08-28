package main

import (
	"html/template"
	"log"
	"net/http"
)

// scans the templates/ folder and parses all .html files there into a *template.Template object
var templates = template.Must(template.ParseGlob("templates/*.html"))

// Combine both template sets
func init() {
	//ParseGlob adds to the existing template set
	//adds all the page templates into that same set
	//This is so that, when rendered, go can combine them (layout + content blocks)
	template.Must(templates.ParseGlob("pages/*.html"))
}

func renderTemplate(w http.ResponseWriter, page string) {
	log.Println("page is ", page)
	//Render layout.html as the base, and pass data into it.
	err := templates.ExecuteTemplate(w, page, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// Routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "index.html")
		log.Println("routed to /")
	})
	http.HandleFunc("/practices", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "practices.html")
	})
	http.HandleFunc("/schedule", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "schedule.html")
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
