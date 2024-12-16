package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type user struct {
	Nome  string
	Email string
}

func main() {

	templates = template.Must(template.ParseGlob("19-HTTP/*.html"))

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		usr := user{
			"Marcos",
			"mail@mail.com",
		}

		templates.ExecuteTemplate(w, "home.html", usr)
	})

	log.Fatal(http.ListenAndServe(":5500", nil))

}
