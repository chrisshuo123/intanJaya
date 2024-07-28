package main

import (
	"html/template"
	"log"
	"net/http"
)

// tpl is a pointer package from Tpl
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/process", processor)
	if err := http.ListenAndServe(":3030", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// Web is about making certain amount of requests
func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func processor(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	email := r.FormValue("emailUser")
	password := r.FormValue("passwordUser")

	d := struct {
		Email    string
		Password string
	}{
		Email:    email,
		Password: password,
	}

	tpl.ExecuteTemplate(w, "processor.html", d)
}
