package parser

import (
    "bufio"
    "strings"
)

type functionScanner struct{
    items []string
}

func (f *functionScanner) match(line string) bool {
    if(len(line) >= 6){ 
        //todo: check for uppercase rune.
        if(line[0:5] == "func "){
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