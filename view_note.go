package main

import (
	"fmt"
	"net/http"
	"strings"
)

func HandleViewNote(w http.ResponseWriter, r *http.Request) {
	file_name := r.URL.Path[7:]
	note, _ := LoadNote(file_name)

	view, _ := Structures.ReadFile("structures/view.html")
	html := Structure(note.Title, []string{"view"}, view)

	html = strings.Replace(html, "{{title}}", note.Title, 2)
	html = strings.Replace(html, "{{content}}", string(note.Body), 1)
	html = strings.Replace(html, "{{file_name}}", file_name, 2)

	fmt.Fprintf(w, "%s", html)
}
