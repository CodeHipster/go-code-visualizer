package main

import (
	"bufio"
	"github.com/codehipster/go-code-visualizer/formatter"
	"github.com/maelkum/go-code-visualizer/parser"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	var dir string

	if len(os.Args) > 1 {
		dir = os.Args[1]
	} else {

		// get directory where binary is located if no args
		// CWD makes more sense but let's keep existing behaviour as-is

		exe_dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			log.Fatal(err)
		}

		dir = exe_dir
	}

	parsedGoCodeFiles := make([]formatter.ParsedCode, 0)

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
	writer.Flush()
}
