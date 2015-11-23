package parser

import(
	goParser "go/parser"
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

func validSyntax(file []byte) bool{	
    _, err := format.Source(file)    
	if(err == nil){
		return true
	}else{
		return false
		} 
}

func goParseFile(path string) (ast *ast.File){
	fset := token.NewFileSet()
	ast, err := goParser.ParseFile(fset, path, nil, 0)
	check(err)
	return
}


//TODO: assuming the code resides in a src folder. is that always true?
//TODO: what if index is -1?
func distilPackagePath(path string) string {	
    directory := filepath.Dir(path)
    srcIndex := strings.Index(directory,"src")
    return filepath.ToSlash(directory[srcIndex+4:])
}

func parseImportDeclarations(importSpecs []ast.Spec) []string{
	imports := make([]string, len(importSpecs))
	for i,spec := range importSpecs{										
		importSpec := spec.(*ast.ImportSpec)
		importString , _ := strconv.Unquote(importSpec.Path.Value)
		imports[i] = importString
	}
	return imports			
}

func parseTypeDeclarations(typeSpecs []ast.Spec) []string{
	//https://golang.org/ref/spec#Types
	
	types := make([]string,len(typeSpecs))
	
	for i,spec := range typeSpecs{								
		typeSpec := spec.(*ast.TypeSpec)
		switch typeType := typeSpec.Type.(type){
			case *ast.Ident:
				ident := (*ast.Ident)(typeType)
				types[i] = "type " + typeSpec.Name.Name + " " + ident.Name
			case *ast.InterfaceType:
				types[i] = "type " + typeSpec.Name.Name + " interface"
			case *ast.StructType:														
				types[i] = "type " + typeSpec.Name.Name + " struct"		
			case *ast.ArrayType:			
			//todo add array length?											
				types[i] = "type []" + typeSpec.Name.Name	
			case *ast.StarExpr:														
				types[i] = "type *" + typeSpec.Name.Name	
			case *ast.FuncType:									
			//TODO: add signature?					
				types[i] = "type " + typeSpec.Name.Name + " func"	
			case *ast.MapType:	
			//todo: add key and element type?													
				types[i] = "type []" + typeSpec.Name.Name	
			case *ast.ChanType:
			//todo: add channel direction?														
				types[i] = "type " + typeSpec.Name.Name + " chan"				
			default:
				fmt.Printf("unknown type: %s\n",reflect.TypeOf(typeSpec.Type))								
		}								
	}	
	return types
}

func parseVariableDeclarations(variableSpecs []ast.Spec) []string{
	variables := make([]string,len(variableSpecs))
	
	for i,spec := range variableSpecs{								
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
			variables[i] = "var " + name.Name + varType
		}
	}	
	return variables
}


func parseConstDeclarations(ConstSpecs []ast.Spec) []string{	
	consts := make([]string,len(ConstSpecs))
	
	for i,spec := range ConstSpecs{								
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
			consts[i] = "const " + name.Name + constType
		}
	}
	return consts
}

func parseGenericDeclaration(genDecl *ast.GenDecl, pgc *ParsedGoCode){
	switch(genDecl.Tok){
		case token.IMPORT:
			pgc.addImports(parseImportDeclarations(genDecl.Specs))
		case token.TYPE:
			pgc.addTypes(parseTypeDeclarations(genDecl.Specs))	
		case token.VAR:	
			pgc.addVars(parseVariableDeclarations(genDecl.Specs))											
		case token.CONST:
			pgc.addVars(parseConstDeclarations(genDecl.Specs))
	}				
}

func ParseFile(path string) (parsedFile ParsedGoCode){	
	    
    //convert file to byte array.
    fileInBytes, err := ioutil.ReadFile(path)
	check(err) 
	
	if(validSyntax(fileInBytes) == false){
		//log file is invalid
		fmt.Printf("File %s has invalid go syntax", path)
	}
	
	syntaxTree := goParseFile(path)
		
    parsedFile.PackagePath = distilPackagePath(path)
	parsedFile.PackageName = syntaxTree.Name.Name
	
	for _, declaration := range syntaxTree.Decls {
		switch declType := declaration.(type) {
				case *ast.GenDecl:			
					genDecl := (*ast.GenDecl)(declType)
					parseGenericDeclaration(genDecl, &parsedFile)
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