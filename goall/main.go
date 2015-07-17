package main

import (
	"log"
	"os"
	"path"

	"github.com/ifo/goall"
)

func main() {
	// ensure example/site directory exists
	err := os.Mkdir("./example/site/posts", 0755)
	if err != nil {
		log.Printf("Warn: %v\n", err)
	}

	// get all posts
	posts := goall.MakePostsList("./example/posts")

	// parse posts into site/posts
	for _, p := range posts {
		prefix := p[:len(p)-len(path.Ext(p))]
		goall.ParseMarkdown("example/posts/"+p, "example/site/posts/"+prefix+".html")
	}
}
