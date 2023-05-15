package main

import (
	"fmt"
	"os"
	"strings"
)

func UpdateNote(fn string, t string, b string) {
	filepath := fmt.Sprintf("notes/%s.txt", fn)
	os.WriteFile(filepath, []byte(fmt.Sprintf("%s\n%s", t, b)), 0644)
	file, _ := os.ReadFile("titles.yml")
	lines := strings.Split(string(file), "\n")
	for i, line := range lines {
		if strings.Contains(line, fn) {
			lines[i] = fn + ": " + t
		}
	}
	os.WriteFile("titles.yml", []byte(strings.Join(lines, "\n")), 0644)
}

func LoadNote(file_name string) (*Note, error) {
	filename := "notes/" + file_name + ".txt"
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	content := string(file)
	lines := strings.Split(content, "\n")
	Title := lines[0]
	Body := strings.Join(lines[1:], "<br />")

	return &Note{Title, Body}, nil
}
