package goalls

import (
	"io/ioutil"

	"github.com/russross/blackfriday"
)

func ParseMarkdown(inputFile string) ([]byte, error) {
	in, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return nil, err
	}

	return blackfriday.MarkdownCommon(in), nil
}
