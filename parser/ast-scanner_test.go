//test
// import, functions, types, variables, constants, multiple and single declaration.
package parser

import (
    "testing"
    "reflect"
    "path/filepath"
    "github.com/davecgh/go-spew/spew")

func TestParser(t *testing.T) {
    absolutePath,_ := filepath.Abs(".\\test-data\\go-code.src"); 
    parsedGoCode := ParseFile(absolutePath)
    
    spew.Dump(parsedGoCode)
    
    expectedParsedGoCode := parsedCode{}
    expectedParsedGoCode.fileName = "go-code.src"
    expectedParsedGoCode.packagePath = "github.com/codehipster/go-code-visualizer/parser/test-data"
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
        
    spew.Dump(expectedParsedGoCode)
        
    if(reflect.DeepEqual(parsedGoCode, expectedParsedGoCode) == false){
        t.Fail();
    }
    
}