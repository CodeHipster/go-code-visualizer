package parser

import (
    "io/ioutil"
    "bufio"
    "strings"
    "go/format"
    "bytes"
    "fmt"
    "path/filepath"
)
//TODO: validate the file for correct go syntax.
//TODO: Implement something more fancy then check() :)
func check(e error) {
    if e != nil {
        panic(e)
    }
}

type ParsedGoCode struct{
    PackageName string
    PackagePath string
	Functions []string
	Types []string
	Imports []string
	Variables []string
}

type goScanner interface {
    match(string) bool
    scan(*bufio.Scanner)
}

func ParseGoCode(path string) ParsedGoCode {
    fmt.Printf("parsing: %s\n", path)
    
    //convert file to byte array.
    src, err := ioutil.ReadFile(path)
	check(err)  
    
    //check go syntax.
    formatSrc, err := format.Source(src)    
	check(err)   
    
    //register scanners
    is := new(importScanner)
    fs := new(functionScanner)
    ts := new(typeScanner)
    vs := new(variableScanner)
    goScanners := []goScanner{is, fs ,ts ,vs} 
	
    parsedFile := ParsedGoCode{}
    
	scanner := bufio.NewScanner(bytes.NewReader(formatSrc))
	scanner.Split(bufio.ScanLines)
    
    //Get package.
    //TODO: first line does not have to be the package declaration
    //Scan for the first encounter of "package " which is not commented.
    scanner.Scan()
    parsedFile.PackageName = scanner.Text()[8:]
    
    //Get path.
    //TODO: what if index is -1?
    directory := filepath.Dir(path)
    index := strings.Index(directory,"src")
    parsedFile.PackagePath = directory[index+4:] 
    
	for scanner.Scan() {
        for _,gs := range goScanners{   
            line := scanner.Text()
            if(gs.match(line)){
                gs.scan(scanner)
                continue
            }  
        }
	}
	
    parsedFile.Imports = is.items
    parsedFile.Functions = fs.items
    parsedFile.Types = ts.items
    parsedFile.Variables = vs.items
        
    return parsedFile
}
