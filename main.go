package main

import (
	"flag"
	"log"
	"os"
	"path"

	"github.com/ifo/goall/goalls"
)

func main() {
	templatesDir := flag.String(
		"templates-dir",
		"./_templates",
		"directory containing templates")

	// setup
	flag.Parse()
	goalls.SetupTemplates(*templatesDir)

	// ensure example/site directory exists
	err := os.MkdirAll("./_site/posts", 0755)
	if err != nil {
		log.Printf("Warn: %v\n", err)
	}

	errIndex := CreateIndexPage(".", "./_posts")
	if errIndex != nil {
		log.Panicln("Could not assemble index page", errIndex)
	}

	CreatePosts("./_posts", "./_site/posts")
}

func CreateIndexPage(rootDir, postsDir string) error {
	index, err := goalls.AssemblePage(
		"index.html",
		goalls.CreateIndex(rootDir, postsDir))
	if err != nil {
		return err
	}

	goalls.OverwriteFile("_site/index.html", index)
	return nil
}

func CreatePosts(postsDir, siteDir string) {
	posts := goalls.MakePostsList(postsDir)

	for _, p := range posts {
		prefix := p[:len(p)-len(path.Ext(p))]
		b, err := goalls.ParseMarkdown(postsDir + "/" + p)
		if err != nil {
			log.Println(err)
			continue
		}
		post, err := goalls.AssembleTemplate("post.html", goalls.TemplateHTML(b))
		if err != nil {
			log.Println(err)
			continue
		}
		goalls.WriteFile(siteDir+"/"+prefix+".html", post)
	}
}
