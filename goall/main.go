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

	// get all posts
	posts := goall.MakePostsList("./posts")

	// parse posts into site/posts
	for _, p := range posts {
		prefix := p[:len(p)-len(path.Ext(p))]
		b, err := goall.ParseMarkdown("posts/" + p)
		if err != nil {
			log.Println(err)
		} else {
			goall.WriteFile("site/posts/"+prefix+".html", b)
		}
	}
}
