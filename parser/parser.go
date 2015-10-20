package parser

import (
    "fmt"
    "os"
    "path/filepath"
    "bufio"
    "strings"
)

//TODO:Implement something more fancy :) 
func check(e error) {
    if e != nil {
        panic(e)
    }
}

type ParsedGoFile struct{
	functions []string
	types []string
	imports []string
	variables []string
}

type goScanner interface {
    match(string) bool
    scan(*bufio.Scanner)
}

type scannedItems struct{
    items []string
}

type functionScanner scannedItems
type typeScanner scannedItems
type importScanner scannedItems
type variableScanner scannedItems

func (f *functionScanner) match(line string) bool {
      //Check if line starts with func?
    if(len(line) >= 6){ 
        //todo: check for uppercase rune.
        if(line[0:5] == "func "){
            return true
        }
    }
    return false   
}

func (f *functionScanner) scan(scanner *bufio.Scanner) {  
    //Check for first occurance of '{'
    line := scanner.Text()
    index := strings.Index(line,"{")
    if(index != -1){
        f.items = append(f.items,line[0:index])
    }
}

func (t *typeScanner) match(line string) bool {
    //Check if line starts with func?
    if(len(line) >= 6){ 
        //todo: check for uppercase rune.
        if(line[0:5] == "type "){
            return true
        }
    }
    return false
}

func (t *typeScanner) scan(scanner *bufio.Scanner) {
    t.items = append(t.items,scanner.Text())
}

func (v *variableScanner) match(line string) bool {
     //Check if line starts with func?
    if(len(line) >= 5){ 
        //todo: check for uppercase rune.
        if(line[0:4] == "var "){
            return true
        }
    }   
    return false
}
func (v *variableScanner) scan(scanner *bufio.Scanner) {
     v.items = append(v.items,scanner.Text())  
}

func (i *importScanner) match(line string) bool {
    if(len(line) >= 6){ 
        //todo: check for uppercase rune.
        if(line[0:6] == "import"){
            return true
        }
    } 
    return false
}

//https://golang.org/ref/spec#Import_declarations
func (i *importScanner) scan(scanner *bufio.Scanner) (moreData bool){
    //check for a bracket. that indicates a block with imports.
    line := scanner.Text()
    index := strings.Index(line,"(")
    if(index == -1){
        //take the word after import.
        i.items = append(i.items,line[7:])
        return
    }
    //else scan next lines until we get a closing bracket.
    for{
        scanner.Scan()
        bracketIndex := strings.Index(scanner.Text(), ")")
        if(bracketIndex != -1){
            if(bracketIndex > 0){
                i.items = append(i.items,scanner.Text()[0:bracketIndex])
            }
            return
        }else{
            i.items = append(i.items, scanner.Text())
        }
    }
}


func ParseGoCode(path string) ParsedGoFile {
	is := importScanner{}
    fs := functionScanner{}
    ts := typeScanner{}
    vs := variableScanner{}

	//Read file contents.
	f, err := os.Open(path)
	defer f.Close()
	
	check(err)        
	
	fmt.Printf("%s\n", filepath.Base(path))
	
	scanner := bufio.NewScanner(bufio.NewReader(f))
	
	scanner.Split(bufio.ScanLines)
	// Validate the input
	for scanner.Scan() {
		line := scanner.Text()
		if(is.match(line)){
			is.scan(scanner)
		}        
		if(fs.match(line)){
			fs.scan(scanner)
		}
		if(ts.match(line)){
			ts.scan(scanner)
		}
		if(vs.match(line)){
			vs.scan(scanner)
		}
	}
	
    parsedFile := ParseGoCode{imports: is.items, functions: fs.items, types:ts.items,variables:vs.items}
    
	for _,element := range is.items {        
		fmt.Printf("%s\n",element)
	}
	for _,element := range fs.items {        
		fmt.Printf("%s\n",element)
	}
	for _,element := range ts.items {        
		fmt.Printf("%s\n",element)
	}
	for _,element := range vs.items {        
		fmt.Printf("%s\n",element)
	}
	
	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s", err)
	}
        
}
