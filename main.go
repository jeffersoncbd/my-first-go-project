package main

import (
	"fmt"
	"log"
	"net/http"

	"os"
)

func main() {
	if _, err := os.Stat("titles.yml"); err != nil {
		os.WriteFile("titles.yml", []byte(""), 0644)
	}

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/css/", HandleCss)

	http.HandleFunc("/notes/", HandleViewNote)

	http.HandleFunc("/new/", HandleNewNoteForm)
	http.HandleFunc("/save", HandleSaveNote)

	http.HandleFunc("/edit/", HandleEditNoteForm)
	http.HandleFunc("/update/", HandleUpdateNote)

	http.HandleFunc("/delete/", HandleDeleteNote)

	http.HandleFunc("/", HandleHome)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HandleCss(w http.ResponseWriter, r *http.Request) {
	cssFileName := r.URL.Path[5:]
	content, _ := os.ReadFile("structures/" + cssFileName)

	fmt.Fprintf(w, "%s", content)
}
