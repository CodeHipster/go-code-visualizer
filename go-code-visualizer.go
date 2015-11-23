package main

import (
	"bufio"
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


	//Create/overwrite a file
	cvFile, err := os.Create(dir + "/dot-visual.cv")
	check(err)
	defer cvFile.Close()

	writer := bufio.NewWriter(cvFile)	

	//walk the filesystem.
	walkFunc := func(path string, info os.FileInfo, err error) error {
		//Skip .git directory.
		if info.IsDir() && info.Name() == ".git" {
			return filepath.SkipDir
		}
		
		//Parse if file is .go file.
		extension := filepath.Ext(path)
		if strings.ToLower(extension) == ".go" {
			parsedFile := parser.ParseFile(path)
			writer.WriteString(parsedFile.ToString())
			writer.Flush()
		}

		return nil
	}
	
	filepath.Walk(dir, walkFunc)
	
	cvFile.Sync()
}
