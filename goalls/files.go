package goalls

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func GetLinksPages(dir string) ([]string, error) {
	out := []string{}

	dirs, err := ioutil.ReadDir(dir)
	if err != nil {
		return out, err
	}

	for _, d := range dirs {
		// ignore all files begining with '_'
		if d.Name()[0] != '_' {
			out = append(out, d.Name())
		}
	}

	return out, nil
}

func MakePostsList(postsDir string) []string {
	dirs, err := ioutil.ReadDir(postsDir)
	if err != nil {
		log.Panicln(err)
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

func MakePostsNames(posts []string) []string {
	out := []string{}
	for _, p := range posts {
		parts := SeparateFileType(p)
		parts[1] = ".html"
		out = append(out, strings.Join(parts, ""))
	}
	return out
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

func SeparateFileType(file string) []string {
	for i := len(file) - 1; i >= 0; i-- {
		if file[i] == '.' {
			return []string{file[:i], file[i:]}
		}
	}
	return []string{file}
}

func ContainsStem(haystack []string, needle string) bool {
	for _, s := range haystack {
		if strings.ToLower(SeparateFileType(s)[0]) ==
			strings.ToLower(SeparateFileType(needle)[0]) {
			return true
		}
	}
	return false
}

func CleanDir(dir string) error {
	// TODO do not allow certain dirs to be cleaned, e.g. _templates
	return os.RemoveAll(dir)
}
