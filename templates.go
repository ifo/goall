package goall

import (
	"bufio"
	"bytes"
	"html/template"
	"log"
)

type Index struct {
	Links []string
	Posts []string
}

// TODO handle errors
func CreateIndex(rootDir, postsDir string) Index {
	links := GetLinksPages(rootDir)
	posts := MakePostsNames(MakePostsList(postsDir))
	return Index{Links: links, Posts: posts}
}

func AssembleIndex(index Index) ([]byte, error) {
	page, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln("template.ParseFiles error", err)
	}

	out := new(bytes.Buffer)
	w := bufio.NewWriter(out)
	err = page.Execute(w, index)
	if err != nil {
		log.Fatalln("template.Execute error", err)
	}
	err = w.Flush()
	if err != nil {
		log.Fatalln("writer Flush error", err)
	}

	// TODO actually return errors
	return out.Bytes(), nil
}

func AssemblePost(post []byte) ([]byte, error) {
	page, err := template.ParseFiles("post.html")
	if err != nil {
		log.Fatalln("template.ParseFiles", err)
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

	// TODO actually return errors
	return out.Bytes(), nil
}
