package goalls

import (
	"bufio"
	"bytes"
	"html/template"
	"log"
)

var templatesDir = "./_templates"

type Index struct {
	Links []string
	Posts []string
}

func SetupTemplates(dir string) {
	if dir != "" {
		templatesDir = dir
	}
}

// TODO handle errors
func CreateIndex(rootDir, postsDir string) Index {
	links, err := GetLinksPages(rootDir)
	if err != nil {
		log.Panicln("GetLinksPages error", err)
	}
	posts := MakePostsNames(MakePostsList(postsDir))
	return Index{Links: links, Posts: posts}
}

func AssemblePage(loc string, info interface{}) ([]byte, error) {
	page, errTemplate := template.ParseFiles(loc)
	if errTemplate != nil {
		log.Println("template.ParseFiles error")
		return nil, errTemplate
	}

	out := new(bytes.Buffer)
	w := bufio.NewWriter(out)
	errExecute := page.Execute(w, info)

	if errExecute != nil {
		log.Println("template.Execute error")
		return nil, errExecute
	}
	errFlush := w.Flush()
	if errFlush != nil {
		log.Println("writer Flush error")
		return nil, errFlush
	}

	return out.Bytes(), nil
}

func AssembleTemplate(loc string, info interface{}) ([]byte, error) {
	return AssemblePage(templatesDir+"/"+loc, info)
}

func TemplateHTML(tmpl []byte) template.HTML {
	return template.HTML(tmpl)
}
