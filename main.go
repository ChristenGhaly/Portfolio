package main

import (
	"encoding/xml"
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

type UrlSet struct {
	XMLName xml.Name `xml:"urlset"`
	Xmlns   string   `xml:"xmlns,attr"`
	Urls    []Url    `xml:"url"`
}

type Url struct {
	Loc        string  `xml:"loc"`
	Changefreq string  `xml:"changefreq"`
	Priority   float32 `xml:"priority"`
}

func sitemapHandler(w http.ResponseWriter, r *http.Request) {
	urls := []Url{
		{Loc: "https://christenghalyportfolio.onrender.com/", Changefreq: "monthly", Priority: 1.0},
		{Loc: "https://christenghalyportfolio.onrender.com/about", Changefreq: "monthly", Priority: 0.8},
		{Loc: "https://christenghalyportfolio.onrender.com/projects", Changefreq: "monthly", Priority: 0.8},
		{Loc: "https://christenghalyportfolio.onrender.com/skills", Changefreq: "monthly", Priority: 0.7},
		{Loc: "https://christenghalyportfolio.onrender.com/experiences", Changefreq: "monthly", Priority: 0.7},
		{Loc: "https://christenghalyportfolio.onrender.com/contact", Changefreq: "monthly", Priority: 0.7},
	}

	urlSet := UrlSet{
		Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9",
		Urls:  urls,
	}

	w.Header().Set("Content-Type", "application/xml")
	xml.NewEncoder(w).Encode(urlSet)
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

	http.HandleFunc("/sitemap.xml", sitemapHandler)

	err := http.ListenAndServe("0.0.0.0:"+port, nil)
	if err != nil {
		log.Fatal(err)
		fmt.Println("Server error:", err)
	}
}