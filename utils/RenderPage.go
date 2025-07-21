package utils

import (
	"garg/resources"
	"html/template"
	"io/fs"
	"net/http"
	"path/filepath"
)

func RenderPage(w http.ResponseWriter, pageFile string, data any) {
	pageFile = filepath.Clean(pageFile) + ".html"

	var tmpl *template.Template
	var err error
	if IsDevMode() {
		tmpl, err = template.ParseGlob("resources/templates/*.templ")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl, err = tmpl.ParseFiles(filepath.Join("resources", "pages", pageFile))
		if err != nil {
			http.Error(w, "Page template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		rootFS, err := fs.Sub(resources.Resources, ".")

		if err != nil {
			http.Error(w, "FS error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl, err = template.ParseFS(rootFS, "templates/*.templ", "pages/"+pageFile)
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	err = tmpl.ExecuteTemplate(w, "head", data)
	if err != nil {
		http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
	}
}
