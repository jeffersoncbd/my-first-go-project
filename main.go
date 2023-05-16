package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"os"
)

//go:embed static/*
var statics embed.FS

func main() {
	if _, err := os.Stat("titles.yml"); err != nil {
		println("creating \"titles.yml\" file...")
		os.WriteFile("titles.yml", []byte(""), 0644)
	}
	if _, err := os.Stat("notes"); err != nil {
		println("creating \"notes\" folder...")
		os.Mkdir("notes", 0760)
	}

	var staticFS = fs.FS(statics)
	contents, _ := fs.Sub(staticFS, "static")
	fs := http.FileServer(http.FS(contents))

	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/css/", HandleCss)

	http.HandleFunc("/notes/", HandleViewNote)

	http.HandleFunc("/new/", HandleNewNoteForm)
	http.HandleFunc("/save", HandleSaveNote)

	http.HandleFunc("/edit/", HandleEditNoteForm)
	http.HandleFunc("/update/", HandleUpdateNote)

	http.HandleFunc("/delete/", HandleDeleteNote)

	http.HandleFunc("/", HandleHome)

	println("server started at :8080 port!")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HandleCss(w http.ResponseWriter, r *http.Request) {
	cssFileName := r.URL.Path[5:]
	content, _ := os.ReadFile("structures/" + cssFileName)

	fmt.Fprintf(w, "%s", content)
}
