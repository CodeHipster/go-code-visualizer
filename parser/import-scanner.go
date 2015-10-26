package parser

import (
    "bufio"
    "strings"
    "strconv"
)

type importScanner struct{
    items []string
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
func (i *importScanner) scan(scanner *bufio.Scanner) {
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
        line = scanner.Text()
        bracketIndex := strings.Index(line, ")")
        if(bracketIndex != -1){
            if(bracketIndex > 0){
                pkg := line[9:]
                pkg = strings.TrimSpace(pkg)
                pkg,_ = strconv.Unquote(pkg)//TODO: catch error
                i.items = append(i.items,pkg)
            }
            return
        }else{
            pkg := line[0:]
            pkg = strings.TrimSpace(pkg)
            pkg,_ = strconv.Unquote(pkg) //TODO: error handing.
            i.items = append(i.items, pkg)
        }
    }
}
