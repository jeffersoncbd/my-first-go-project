package main

import (
	"net/http"
	"os"
	"strings"
)

func HandleDeleteNote(w http.ResponseWriter, r *http.Request) {
	file_name := r.URL.Path[8:]
	os.Remove("notes/" + file_name + ".txt")
	new_file := []string{}
	file, _ := os.ReadFile("titles.yml")
	lines := strings.Split(string(file), "\n")
	for _, line := range lines {
		if strings.Contains(line, file_name) {
			continue
		}
		new_file = append(new_file, line)
	}
	os.WriteFile("titles.yml", []byte(strings.Join(new_file, "\n")), 0644)
	http.Redirect(w, r, "/", http.StatusFound)
}
