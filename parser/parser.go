package parser

import (
    "fmt"
)

//TODO: Implement something more fancy then check() :)
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func ParseGoCode(path string) ParsedGoCode {
    fmt.Printf("parsing: %s\n", path)  
	
    parsedFile := ParseFile(path)      
        
    return parsedFile
}
