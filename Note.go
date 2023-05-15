package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Note struct {
	Title string
	Body  string
}

func (n *Note) save() string {
	time := time.Now().UnixMilli()
	random := rand.Int63n(8999) + 1000
	filename := fmt.Sprintf("%v-%v", time, random)
	filepath := "notes/" + filename + ".txt"
	os.WriteFile(filepath, []byte(n.Title+"\n"+n.Body), 0644)
	file, _ := os.OpenFile("titles.yml", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	file.Write([]byte(filename + ": " + n.Title + "\n"))
	file.Close()
	return filename
}
