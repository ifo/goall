package goall

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/russross/blackfriday"
)

func MakePostsList(postsDir string) []string {
	// get dir list
	dirs, err := ioutil.ReadDir(postsDir)
	if err != nil {
		log.Fatalln(err)
	}

	out := []string{}
	for _, d := range dirs {
		if strings.HasSuffix(d.Name(), ".md") ||
			strings.HasSuffix(d.Name(), ".markdown") {
			out = append(out, d.Name())
		}
	}

	return out
}

func ParseMarkdown(inputFile string) ([]byte, error) {
	in, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return nil, err
	}

	return blackfriday.MarkdownCommon(in), nil
}

func WriteFile(dest string, info []byte) error {
	_, err := os.Stat(dest)
	if os.IsNotExist(err) {
		return OverwriteFile(dest, info)
	}
	return err
}

func OverwriteFile(dest string, info []byte) error {
	outFile, errCreate := os.Create(dest)
	if errCreate != nil {
		return errCreate
	}
	defer outFile.Close()

	w := bufio.NewWriter(outFile)
	_, errWrite := w.Write(info)
	w.Flush()
	return errWrite
}
