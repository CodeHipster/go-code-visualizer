package main

import (
	"fmt"
	"github.com/thijsoostdam/go-code-visualizer/parser"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//todo: skip hidden directories.
//http://grokbase.com/t/gg/golang-nuts/144va1n8w5/go-nuts-how-do-check-if-file-or-directory-is-hidden-under-windows
//Check if syscall hidden attribute works for windows and unix systems.

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	//Get current directory
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	//walk the filesystem.
	walkFunc := func(path string, info os.FileInfo, err error) error {
		//Check if it is .git directory (improve to all hidden files and folders.)
		if info.IsDir() && info.Name() == ".git" {
			return filepath.SkipDir
		}
		//Check extension
		extension := filepath.Ext(path)
		if strings.ToLower(extension) != ".go" {
			return nil
		}

		parsedFile := parser.ParseGoCode(path)

		fmt.Printf("%+v\n", parsedFile)

		return nil
	}
	filepath.Walk(dir, walkFunc)

	//Create/overwrite a file
	f, err := os.Create(dir + "/dot-visual.cv")
	check(err)

	defer f.Close()

	f.Sync()
}
