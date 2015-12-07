package main

import (
	"bufio"
	"github.com/thijsoostdam/go-code-visualizer/parser"
	"github.com/thijsoostdam/go-code-visualizer/formatter"
	"log"
	"os"
	"path/filepath"
	"strings"
	"fmt"
)

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

	parsedGoCodeFiles := make([]formatter.ParsedCode,0)

	//walk the filesystem.
	walkFunc := func(path string, info os.FileInfo, err error) error {
		//Skip .git directory.
		if info.IsDir() && info.Name() == ".git" {
			return filepath.SkipDir
		}
		
		//Parse if file is .go file.
		extension := filepath.Ext(path)
		if strings.ToLower(extension) == ".go" {
			parsedGoCode := parser.ParseFile(path)
			parsedGoCodeFiles = append(parsedGoCodeFiles, parsedGoCode)
		}

		return nil
	}
	
	filepath.Walk(dir, walkFunc)
	
	dotGraph := formatter.GenerateDotGraph(parsedGoCodeFiles)
	//Create/overwrite a file
	cvFile, err := os.Create(dir + "/dot-visual.gv")
	check(err)
	defer cvFile.Close()
	
	writer := bufio.NewWriter(cvFile)	
	
	writer.WriteString(dotGraph)
	
	cvFile.Sync()
}
