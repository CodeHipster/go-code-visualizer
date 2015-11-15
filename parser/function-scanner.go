package parser

import (
    "bufio"
    "strings"
    "unicode"  
    "regexp"  
)

type functionScanner struct{
    items []string
}

func (f *functionScanner) match(line string) bool {
    if(len(line) >= 6){ 
        //todo: check for uppercase rune.
        if(line[0:5] == "func "){
            if(strings.IsUpperCase(line[6:7])
            return true
        }
    }
    return false   
}

func (f *functionScanner) scan(scanner *bufio.Scanner) { 
    line := scanner.Text()
    index := strings.Index(line,"{")
    if(index != -1){
        f.items = append(f.items,line[5:index])
    }
}

func extractFuncName(string line){
    //https://golang.org/ref/spec#Function_declarations
    //https://golang.org/ref/spec#Method_declarations
    //https://play.golang.org/p/7ccWVkM2kc
    //check if it is a method or function.
	reFunc := `(^func\b)`
	reReceiver := `(\([^\)]+\))?`
	reMethodName := `(\b\w+\b)`
	reParameters := `(\([^\)]*\))`
	reResult := `((?:\([^\)]+\))?|(?:\w+)?)`
	reWhiteSpace := `\s*`
	
    re := regexp.MustCompile(reFunc+reWhiteSpace+reReceiver+reWhiteSpace+reMethodName+reParameters+reWhiteSpace+reResult)
	
    //TODO:
    extract methodName  
}