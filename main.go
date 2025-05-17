package main

import (
	"fmt"
	"html/template"
	"net/http"
	// "gopkg.in/gomail.v2"
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

// func sendEmail(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
//         http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
//         return
//     }

// 	r.ParseForm()
// 	name := r.FormValue("full-name")
// 	email := r.FormValue("email")
// 	subject := r.FormValue("subject")
// 	message := r.FormValue("message")

// 	msg := gomail.NewMessage()
// 	msg.SetHeader("From", "christen.srghaly@yahoo.com")
// 	msg.SetHeader("To", "christen.srghaly@yahoo.com")
// 	msg.SetHeader("Subject", subject)
// 	msg.SetBody("text/plain", fmt.Sprintf("From: %s <%s>\n\n%s", name, email, message))

// 	d := gomail.NewDialer("smtp.gmail.com", 587, "christen.srghaly@yahoo.com", "your-password")

// 	if err := d.DialAndSend(msg); err != nil {
// 		http.Error(w, "Failed to send email: "+err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Write([]byte("Message sent successfully!"))
// }

func main() {
	fmt.Println("Server running at http://localhost:8080")

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { renderTemplate(w, "index", PageData{Breadcrumb: []string{"Home"}})})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) { renderTemplate(w, "about", PageData{Breadcrumb: []string{"Home", "About"}})})

	http.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) { renderTemplate(w, "projects", PageData{Breadcrumb: []string{"Home", "Projects"}})})

	http.HandleFunc("/skills", func(w http.ResponseWriter, r *http.Request) { renderTemplate(w, "skills", PageData{Breadcrumb: []string{"Home", "Technical Skills"}})})

	http.HandleFunc("/experiences", func(w http.ResponseWriter, r *http.Request) { renderTemplate(w, "experiences", PageData{Breadcrumb: []string{"Home", "Experiences"}})})
	
	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) { renderTemplate(w, "contact", PageData{Breadcrumb: []string{"Home", "Contact"}})})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}