package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	home, _ := Structures.ReadFile("structures/home.html")
	html := Structure("Home", []string{"home"}, home)

	content, _ := os.ReadFile("titles.yml")
	lines := strings.Split(string(content), "\n")

	titles := []string{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, ": ")
		html := fmt.Sprintf("<a href=\"/notes/%s\"><li>%s</li></a>", parts[0], parts[1])
		titles = append(titles, html)
	}

	html = strings.Replace(html, "{{notes}}", strings.Join(titles, ""), 1)

	fmt.Fprintf(w, "%s", html)
}
