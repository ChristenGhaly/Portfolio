package main

import (
	"fmt"
	"html/template"
	"net/http"
	"log"
	"os"
)

type PageData struct {
	Breadcrumb []string
}

func renderTemplate(w http.ResponseWriter, tmpl string, data PageData) {
	t, err := template.ParseFiles(
        "templates/base.html",
        "templates/partials/header.html",
        "templates/partials/footer.html",
        "templates/" + tmpl + ".html",
    )
	if err != nil {
		http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
		fmt.Println("Template error:", err)
		return
	}
	
	err = t.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		http.Error(w, "Execution error: "+err.Error(), http.StatusInternalServerError)
		fmt.Println("Execution error:", err)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server running on port %s\n", port)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { renderTemplate(w, "index", PageData{Breadcrumb: []string{"Home"}})})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) { renderTemplate(w, "about", PageData{Breadcrumb: []string{"Home", "About"}})})

	http.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) { renderTemplate(w, "projects", PageData{Breadcrumb: []string{"Home", "Projects"}})})

	http.HandleFunc("/skills", func(w http.ResponseWriter, r *http.Request) { renderTemplate(w, "skills", PageData{Breadcrumb: []string{"Home", "Technical Skills"}})})

	http.HandleFunc("/experiences", func(w http.ResponseWriter, r *http.Request) { renderTemplate(w, "experiences", PageData{Breadcrumb: []string{"Home", "Experiences"}})})
	
	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) { renderTemplate(w, "contact", PageData{Breadcrumb: []string{"Home", "Contact"}})})

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
		fmt.Println("Server error:", err)
	}
}