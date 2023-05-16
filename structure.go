package main

import (
	"embed"
	"strings"
)

//go:embed structures/*
var Structures embed.FS

func Structure(title string, css_list []string, body []byte) string {
	html_content, _ := Structures.ReadFile("structures/root.html")
	html := string(html_content)

	css_tags := ""
	for _, css := range css_list {
		css_tags = css_tags + "<link rel=\"stylesheet\" href=\"/static/" + css + ".css\">"
	}

	html = strings.Replace(html, "{{title}}", title, 1)
	html = strings.Replace(html, "{{css}}", css_tags, 1)
	html = strings.Replace(html, "{{body}}", string(body), 1)

	return html
}
