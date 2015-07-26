package goall

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"

	"github.com/knieriem/markdown"
)

func MakePostsList(postsDir string) []string {
	// get dir list
	dirs, err := ioutil.ReadDir(postsDir)
	if err != nil {
		log.Fatalln(err)
	}

	out := []string{}
	for _, d := range dirs {
		out = append(out, d.Name())
	}

	return out
}

// parse a markdown file and save it as a post, unless it exists
func ParseMarkdown(input, output string) {
	if _, err := os.Stat(output); os.IsNotExist(err) {
		p := markdown.NewParser(&markdown.Extensions{Smart: true})

		inFile, err := os.Open(input)
		if err != nil {
			panic(err)
		}
		defer inFile.Close()

		outFile, err := os.Create(output)
		if err != nil {
			panic(err)
		}
		defer outFile.Close()

		w := bufio.NewWriter(outFile)

		p.Markdown(inFile, markdown.ToHTML(w))
		w.Flush()
	}
}
