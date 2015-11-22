package parser

import(
	"go/parser"
	"go/token"
	"go/format"
	"fmt"
	"reflect"
	"go/ast"
    "path/filepath"
    "strconv"
	"strings"
    "io/ioutil"
	)
	
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func ParseFile(path string) (parsedFile ParsedGoCode){	
	    
    //convert file to byte array.
    fileInBytes, err := ioutil.ReadFile(path)
	check(err) 
	
	//check go syntax.
    _, err = format.Source(fileInBytes)    
	check(err) 
	
	fset := token.NewFileSet() // positions are relative to fset

	f, err := parser.ParseFile(fset, path, nil, 0)
	if err != nil {
		fmt.Println(err)
		return 
	}	
	
    //Get path.
	//TODO: assuming the code resides in a src folder. is that always true?
    //TODO: what if index is -1?
    directory := filepath.Dir(path)
    index := strings.Index(directory,"src")
    parsedFile.PackagePath = directory[index+4:] 

	parsedFile.PackageName = f.Name.Name
	
	for _, declaration := range f.Decls {
		switch declType := declaration.(type) {
				case *ast.GenDecl:			
					genDecl := (*ast.GenDecl)(declType)
					switch(genDecl.Tok){
						case token.IMPORT: //test with 1 or multiple import statements.
							for _,spec := range genDecl.Specs{								
								importSpec := spec.(*ast.ImportSpec)
								importString , _ := strconv.Unquote(importSpec.Path.Value)
								(&parsedFile).addImport(importString)
							}
							break;
						case token.TYPE:
						//https://golang.org/ref/spec#Types
							for _,spec := range genDecl.Specs{								
								typeSpec := spec.(*ast.TypeSpec)
								switch typeType := typeSpec.Type.(type){
									case *ast.Ident:
										ident := (*ast.Ident)(typeType)
										(&parsedFile).addType("type " + typeSpec.Name.Name + " " + ident.Name)
									case *ast.InterfaceType:
										(&parsedFile).addType("type " + typeSpec.Name.Name + " interface")
									case *ast.StructType:														
										(&parsedFile).addType("type " + typeSpec.Name.Name + " struct")		
									case *ast.ArrayType:			
									//todo add array length?											
										(&parsedFile).addType("type []" + typeSpec.Name.Name)	
									case *ast.StarExpr:														
										(&parsedFile).addType("type *" + typeSpec.Name.Name)	
									case *ast.FuncType:									
									//TODO: add signature?					
										(&parsedFile).addType("type " + typeSpec.Name.Name + " func")	
									case *ast.MapType:	
									//todo: add key and element type?													
										(&parsedFile).addType("type []" + typeSpec.Name.Name)	
									case *ast.ChanType:
									//todo: add channel direction?														
										(&parsedFile).addType("type " + typeSpec.Name.Name + " chan")					
									default:
										fmt.Printf("unknown type: %s\n",reflect.TypeOf(typeSpec.Type))								
								}								
							}
							break;	
						case token.VAR:	
							for _,spec := range genDecl.Specs{								
								varSpec := spec.(*ast.ValueSpec)
								var varType string
								switch typeVar := varSpec.Type.(type){ //check for nil?
									case *ast.Ident:									
										ident := (*ast.Ident)(typeVar)
										varType = ident.Name
									break;
									default:
										fmt.Printf("unknown varType: %s\n",reflect.TypeOf(varSpec.Type))
								}
								for _,name := range varSpec.Names{
									(&parsedFile).addVar("var " + name.Name + varType)
								}
							}
							break;											
						case token.CONST:
							for _,spec := range genDecl.Specs{								
								constSpec := spec.(*ast.ValueSpec)
								var constType string
								switch typeVar := constSpec.Type.(type){
									case *ast.Ident:									
										ident := (*ast.Ident)(typeVar)
										constType = ident.Name
									break;
									default:
										fmt.Printf("unknown varType: %s\n",reflect.TypeOf(constSpec.Type))
								}
								for _,name := range constSpec.Names{
									(&parsedFile).addVar("const " + name.Name + constType)
								}
							}
							break;
					}				
				case *ast.FuncDecl:					
					funcDecl := (*ast.FuncDecl)(declType)
					
					functionString := string(fileInBytes[funcDecl.Pos()-1:funcDecl.Body.Lbrace-1])
					(&parsedFile).addFunc(functionString)
					
				default:
					fmt.Println("unknown declaration")
		}
	}
	
	return parsedFile
}