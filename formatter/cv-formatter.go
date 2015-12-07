package formatter

func GenerateDotGraph(goCode []ParsedCode) string{
	
	dotGraph := ""
	
	for _, parsed := range goCode{
		dotGraph += parsed.ToString()
	}
	
	return dotGraph	
}