package main

import (
	"fmt"
	"net/http"
	"strings"
)

func HandleNewNoteForm(w http.ResponseWriter, r *http.Request) {
	form, _ := Structures.ReadFile("structures/form.html")
	html := Structure("New", []string{"form"}, form)

	html = strings.Replace(html, "{{operation}}", "New", 1)
	html = strings.Replace(html, "{{action}}", "/save", 1)
	html = strings.Replace(html, "{{method}}", "POST", 1)
	html = strings.Replace(html, "{{title}}", "", 1)
	html = strings.Replace(html, "{{body}}", "", 1)
	fmt.Fprint(w, html)
}

func HandleSaveNote(w http.ResponseWriter, r *http.Request) {
	Title := r.FormValue("title")
	Body := r.FormValue("body")
	note := &Note{Title, Body}
	filename := note.save()
	http.Redirect(w, r, "/notes/"+filename, http.StatusFound)
}
