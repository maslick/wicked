package logic

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles("views/edit.html", "views/items.html", "views/view.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	if err := templates.ExecuteTemplate(w, tmpl+".html", p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
