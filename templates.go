package goall

import (
	"bufio"
	"bytes"
	"html/template"
	"log"
)

func AssemblePost(post []byte) ([]byte, error) {
	page, err := template.ParseFiles("post.html")
	if err != nil {
		log.Fatalln("post.html file required", err)
	}

	out := new(bytes.Buffer)
	w := bufio.NewWriter(out)

	err = page.Execute(w, template.HTML(post))
	if err != nil {
		log.Fatalln("template.Execute error", err)
	}
	err = w.Flush()
	if err != nil {
		log.Fatalln("writer Flush error", err)
	}

	// TODO actually pass errors
	return out.Bytes(), nil
}
