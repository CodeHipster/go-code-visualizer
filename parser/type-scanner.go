package parser

import (
    "bufio"
    "strings"
)

type typeScanner struct{
    items []string
}

func (t *typeScanner) match(line string) bool {
    if(len(line) >= 6){ 
        //todo: check for uppercase rune.
        if(line[0:5] == "type "){
            return true
        }
    }
    return false
}

func (t *typeScanner) scan(scanner *bufio.Scanner) {
    line := scanner.Text()[5:]
    index := strings.Index(line,"{")
    if(index != -1){
        t.items = append(t.items,line[0:index-1])
    }else{
        t.items = append(t.items,line)
    }
}