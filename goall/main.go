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

	errIndex := CreateIndexPage(".", "./posts")
	if errIndex != nil {
		log.Fatalln("Could not assemble index page", errIndex)
	}

	CreatePosts("./posts", "./site/posts")
}

func CreateIndexPage(rootDir, postsDir string) error {
	index, err := goall.AssembleIndex(goall.CreateIndex(rootDir, postsDir))
	if err != nil {
		return err
	}

	goall.OverwriteFile("site/index.html", index)
	return nil
}

func CreatePosts(postsDir, siteDir string) {
	posts := goall.MakePostsList(postsDir)

	for _, p := range posts {
		prefix := p[:len(p)-len(path.Ext(p))]
		b, err := goall.ParseMarkdown(postsDir + p)
		if err != nil {
			log.Println(err)
			continue
		}
		post, err := goall.AssemblePost(b)
		if err != nil {
			log.Println(err)
			continue
		}
		goall.WriteFile(siteDir+prefix+".html", post)
	}
}
