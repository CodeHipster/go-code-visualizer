package parser

type ParsedGoCode struct{
    PackagePath string
    PackageName string
	Imports []string
	Types []string
	Variables []string
	Functions []string	
}

func (p ParsedGoCode) addImport(imp string){
	p.Imports = append(p.Imports, imp)
} 

func (p ParsedGoCode) addType(typ string){
	p.Types = append(p.Types, typ)
} 

func (p ParsedGoCode) addVar(variable string){
	p.Variables = append(p.Variables, variable)
} 

func (p ParsedGoCode) addFunc(function string){
	p.Functions = append(p.Functions, function)
} 
