package parser

import "fmt"

type parsedCode struct{
	fileName string
    packagePath string
    packageName string
	imports []string
	types []string
	variables []string
	functions []string	
}

func (p parsedCode) FileName() string{
	return p.fileName
}

func (p parsedCode) PackagePath() string{
	return p.packagePath
}

func (p parsedCode) PackageName() string{
	return p.packageName
}
	
func (p *parsedCode) addImport(imp string){
	p.imports = append(p.imports, imp)
} 

func (p *parsedCode) addImports(imp []string){
	p.imports = append(p.imports, imp...)
} 

func (p parsedCode) Imports() []string{
	return p.imports
}

func (p *parsedCode) addType(typ string){
	p.types = append(p.types, typ)
} 

func (p *parsedCode) addTypes(typ []string){
	p.types = append(p.types, typ...)
} 

func (p parsedCode) Types() []string{
	return p.types
} 

func (p *parsedCode) addVar(variable string){
	p.variables = append(p.variables, variable)
} 

func (p *parsedCode) addVars(variable []string){
	p.variables = append(p.variables, variable...)
} 

func (p parsedCode) Variables() []string{
	return p.variables
} 

func (p *parsedCode) addFunc(function string){
	p.functions = append(p.functions, function)
} 

func (p *parsedCode) addFuncs(function []string){
	p.functions = append(p.functions, function...)
} 

func (p parsedCode) Functions() []string{
	return p.functions
}

func (p parsedCode) ToString() (str string){
	str += fmt.Sprintf("package path: %s\n",p.packagePath)
	str += fmt.Sprintf("package name: %s\n",p.packageName)
	str += fmt.Sprintf("imports:\n")
	for _,imp := range p.imports{		
		str += fmt.Sprintf("%s\n",imp)
	}
	str += fmt.Sprintf("types:\n")
	for _,typ := range p.types{		
		str += fmt.Sprintf("%s\n",typ)
	}
	str += fmt.Sprintf("variables:\n")
	for _,vari := range p.variables{		
		str += fmt.Sprintf("%s\n",vari)
	}
	str += fmt.Sprintf("functions:\n")
	for _,fun := range p.functions{		
		str += fmt.Sprintf("%s\n",fun)
	}
	str += "\n"
	return str
}