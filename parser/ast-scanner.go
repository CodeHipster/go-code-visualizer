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
		imports[i] = strings.ToLower(importString)
	}
	
	return imports			
}

func parseTypeDeclarations(typeSpecs []ast.Spec) []string{
	//https://golang.org/ref/spec#Types
	//make slice with 0 length but with a capacity to hold all types.
	types := make([]string,0,len(typeSpecs))
	
	for _,spec := range typeSpecs{								
		typeSpec := spec.(*ast.TypeSpec)		
		if(ast.IsExported(typeSpec.Name.Name)){
			fmt.Printf("Type exported: %s\n",typeSpec.Name.Name)
		}else{			
			fmt.Printf("Type not exported: %s\n",typeSpec.Name.Name)
			continue
		}
		switch typeType := typeSpec.Type.(type){
			case *ast.Ident:
				ident := (*ast.Ident)(typeType)
				types = append(types,"type " + typeSpec.Name.Name + " " + ident.Name)
			case *ast.InterfaceType:
				types = append(types, "type " + typeSpec.Name.Name + " interface")
			case *ast.StructType:														
				types = append(types, "type " + typeSpec.Name.Name + " struct")		
			case *ast.ArrayType:			
			//todo add array length?											
				types = append(types, "type []" + typeSpec.Name.Name)	
			case *ast.StarExpr:														
				types = append(types, "type *" + typeSpec.Name.Name)	
			case *ast.FuncType:									
			//TODO: add signature?					
				types = append(types, "type " + typeSpec.Name.Name + " func")	
			case *ast.MapType:	
			//todo: add key and element type?													
				types = append(types, "type []" + typeSpec.Name.Name)	
			case *ast.ChanType:
			//todo: add channel direction?														
				types = append(types, "type " + typeSpec.Name.Name + " chan")				
			default:
				fmt.Printf("unknown type: %s\n",reflect.TypeOf(typeSpec.Type))								
		}								
	}	
	return types
}

func parseVariableDeclarations(variableSpecs []ast.Spec, prefix string) []string{
	//make slice with 0 length but with a capacity to hold all variables.
	variables := make([]string,0,len(variableSpecs))
	
	for _,spec := range variableSpecs{								
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
			if(ast.IsExported(name.Name)){
				fmt.Printf("Variable exported: %s\n",name.Name)
				variables = append(variables, prefix + name.Name + varType)
			}else{
				fmt.Printf("Variable not exported: %s\n",name.Name)
			}
		}
	}	
	return variables
}

func parseFunction(funcDecl *ast.FuncDecl, fileInBytes []byte) (string, bool){
	if(ast.IsExported(funcDecl.Name.Name)){
		fmt.Printf("Function exported: %s\n", funcDecl.Name.Name)
		return string(fileInBytes[funcDecl.Pos()-1:funcDecl.Body.Lbrace-1]), true						
	}else{		
		fmt.Printf("Function not exported: %s\n", funcDecl.Name.Name)
	}
	return "", false
}

func ParseFile(path string) (parsedGoCode parsedCode){	
	    
    //convert file to byte array.
    fileInBytes, err := ioutil.ReadFile(path)
	check(err) 
	
	if(validSyntax(fileInBytes) == false){
		fmt.Printf("File %s has invalid go syntax", path)
	}
	
	syntaxTree := goParseFile(path)
		
	parsedGoCode.fileName = filepath.Base(path)
    parsedGoCode.packagePath = strings.ToLower(distilPackagePath(path))
	parsedGoCode.packageName = syntaxTree.Name.Name
	
	for _, declaration := range syntaxTree.Decls {
		switch declType := declaration.(type) {
			case *ast.GenDecl:
				genericDeclaration := (*ast.GenDecl)(declType)
				switch(genericDeclaration.Tok){
					case token.IMPORT:
						parsedGoCode.addImports(parseImportDeclarations(genericDeclaration.Specs))
					case token.TYPE:
						parsedGoCode.addTypes(parseTypeDeclarations(genericDeclaration.Specs))	
					case token.VAR:	
						parsedGoCode.addVars(parseVariableDeclarations(genericDeclaration.Specs, "var "))											
					case token.CONST:
						parsedGoCode.addVars(parseVariableDeclarations(genericDeclaration.Specs, "const "))
					default:
						fmt.Println("unknown generic declaration")
				}
			case *ast.FuncDecl:				
				funcDecl := (*ast.FuncDecl)(declType)	
				funcString, success := parseFunction(funcDecl, fileInBytes)
				if(success) {parsedGoCode.addFunc(funcString)}								
			default:
				fmt.Println("unknown declaration")
		}
	}
	
	return parsedGoCode
}