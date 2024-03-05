package handlers

import (
	"net/http"
	"os"
	"text/template"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	htmlContent, err := os.ReadFile("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Parse the HTML content as a template
	tmpl, err := template.New("index").Parse(string(htmlContent))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the template
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
