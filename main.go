package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	notPermited = []string{".git", "node_modules"}
)

func Includes(arr []string, path string) bool {
	for _, a := range arr {
		if strings.Contains(path, a) {
			return true
		}
	}
	return false
}

func IsExist(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

func verifyDirsRecursively(dir string, word string) error {
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if IsExist(path) && !Includes(notPermited, path) && !info.IsDir() {

			result, err := ioutil.ReadFile(info.Name())
			if err != nil {
				return err
			}
			buf := bytes.Buffer{}
			_, err = buf.Write(result)
			if err != nil {
				return err
			}
			if exists := strings.Contains(string(result), word); exists {

				focusedResult := strings.ReplaceAll(string(result), word, fmt.Sprintf("\x1b[91m%s\x1b[0m", word))
				fmt.Printf("\x1b[92m[PATH]\x1b[0m: \x1b[93m%s\x1b[0m\n%s\n", fmt.Sprintf("%s", path), focusedResult)
			}

		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func main() {
	if len(os.Args) >= 2 {
		word := os.Args[1]
		if err := verifyDirsRecursively(".", word); err != nil {
			log.Println(err.Error())
		}
	} else {
		fmt.Printf("Search is utility of view files of contains word sentences.\nUSAGE: search [word]\n")
	}

}
