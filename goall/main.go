package main

import (
	"log"
	"os"
	"path"

	"github.com/ifo/goall"
)

func main() {
	// ensure example/site directory exists
	err := os.MkdirAll("./site/posts", 0755)
	if err != nil {
		log.Printf("Warn: %v\n", err)
	}

	// create index
	index, err := goall.AssembleIndex(goall.CreateIndex(".", "./posts"))
	if err != nil {
		log.Fatalln("Could not assemble index page", err)
	}

	// index might be updated
	goall.OverwriteFile("site/index.html", index)

	// get all posts
	posts := goall.MakePostsList("./posts")

	// parse posts into site/posts
	for _, p := range posts {
		prefix := p[:len(p)-len(path.Ext(p))]
		b, err := goall.ParseMarkdown("posts/" + p)
		if err != nil {
			log.Println(err)
			continue
		}
		post, err := goall.AssemblePost(b)
		if err != nil {
			log.Println(err)
			continue
		}
		goall.WriteFile("site/posts/"+prefix+".html", post)
	}
}
