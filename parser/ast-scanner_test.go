//test
// import, functions, types, variables, constants, multiple and single declaration.
package parser

import (
    "testing"
    "reflect"
    "github.com/davecgh/go-spew/spew")

func TestParser(t *testing.T) {
    parsedGoCode := ParseFile(".\\test-data\\go-code.src")
    
    spew.Dump(parsedGoCode)
    
    expectedParsedGoCode := parsedCode{}
    expectedParsedGoCode.fileName = "go-code.src"
    expectedParsedGoCode.packagePath = "t-data" //TODO: looks wrong :-/
    expectedParsedGoCode.packageName = "main"
    //expectedParsedGoCode.imports
    expectedParsedGoCode.addVars([]string{
        "var Test int",
        "var TestChan <-chan",
    }) 
    
    expectedParsedGoCode.addTypes([]string{            
        "type Booltype bool",
        "type NumericType uint8",
        "type StringType string",
        "type []ArrayType",
        "type []SliceType",
        "type StructType struct",
        "type *PointerType",
        "type FunctionType func",
        "type InterfaceType interface",
        "type []MapType",
        "type ChannelType <-chan<-",
    })
        
    if(reflect.DeepEqual(parsedGoCode, expectedParsedGoCode) == false){
        t.Fail();
    }
    
}