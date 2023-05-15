package main

import (
	"fmt"
	"net/http"
	"strings"
)

func HandleEditNoteForm(w http.ResponseWriter, r *http.Request) {
	html := Structure("Edit", "form", []string{"form"})

	file_name := r.URL.Path[6:]
	note, _ := LoadNote(file_name)

	body := string(note.Body)
	body = strings.ReplaceAll(body, "<br />", "\n")

	html = strings.Replace(html, "{{operation}}", "Edit", 1)
	html = strings.Replace(html, "{{action}}", "/update/"+file_name, 1)
	html = strings.Replace(html, "{{method}}", "POST", 1)
	html = strings.Replace(html, "{{title}}", note.Title, 1)
	html = strings.Replace(html, "{{body}}", body, 1)
	fmt.Fprint(w, html)
}

func HandleUpdateNote(w http.ResponseWriter, r *http.Request) {
	file_name := r.URL.Path[8:]
	title := r.FormValue("title")
	body := r.FormValue("body")
	UpdateNote(file_name, title, string(body))
	http.Redirect(w, r, "/notes/"+file_name, http.StatusFound)
}
