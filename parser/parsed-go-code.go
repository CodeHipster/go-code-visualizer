package parser

import "fmt"

type ParsedGoCode struct{
    PackagePath string
    PackageName string
	Imports []string
	Types []string
	Variables []string
	Functions []string	
}

func (p *ParsedGoCode) addImport(imp string){
	p.Imports = append(p.Imports, imp)
} 

func (p *ParsedGoCode) addImports(imp []string){
	p.Imports = append(p.Imports, imp...)
} 

func (p *ParsedGoCode) addType(typ string){
	p.Types = append(p.Types, typ)
} 

func (p *ParsedGoCode) addTypes(typ []string){
	p.Types = append(p.Types, typ...)
} 

func (p *ParsedGoCode) addVar(variable string){
	p.Variables = append(p.Variables, variable)
} 

func (p *ParsedGoCode) addVars(variable []string){
	p.Variables = append(p.Variables, variable...)
} 

func (p *ParsedGoCode) addFunc(function string){
	p.Functions = append(p.Functions, function)
} 

func (p *ParsedGoCode) addFuncs(function []string){
	p.Functions = append(p.Functions, function...)
} 

func (p ParsedGoCode) ToString() (str string){
	str += fmt.Sprintf("package path: %s\n",p.PackagePath)
	str += fmt.Sprintf("package name: %s\n",p.PackageName)
	str += fmt.Sprintf("imports:\n")
	for _,imp := range p.Imports{		
		str += fmt.Sprintf("%s\n",imp)
	}
	str += fmt.Sprintf("types:\n")
	for _,typ := range p.Types{		
		str += fmt.Sprintf("%s\n",typ)
	}
	str += fmt.Sprintf("variables:\n")
	for _,vari := range p.Variables{		
		str += fmt.Sprintf("%s\n",vari)
	}
	str += fmt.Sprintf("functions:\n")
	for _,fun := range p.Functions{		
		str += fmt.Sprintf("%s\n",fun)
	}
	return str
}