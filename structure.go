package main

import (
	"os"
	"strings"
)

func Structure(title string, file_name string, css_list []string) string {
	html_content, _ := os.ReadFile("structures/root.html")
	html := string(html_content)

	body_content, _ := os.ReadFile("structures/" + file_name + ".html")
	body := string(body_content)

	css_tags := ""
	for _, css := range css_list {
		css_tags = css_tags + "<link rel=\"stylesheet\" href=\"/static/" + css + ".css\">"
	}

	html = strings.Replace(html, "{{title}}", title, 1)
	html = strings.Replace(html, "{{css}}", css_tags, 1)
	html = strings.Replace(html, "{{body}}", body, 1)

	return html
}
