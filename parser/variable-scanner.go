package parser

import (
    "bufio"
)

type variableScanner struct{
    items []string
}

func (v *variableScanner) match(line string) bool {
    if(len(line) >= 5){ 
        //todo: check for uppercase rune.
        if(line[0:4] == "var "){
            return true
        }
    }   
    return false
}
func (v *variableScanner) scan(scanner *bufio.Scanner) {
     v.items = append(v.items,scanner.Text()[4:])  
}   