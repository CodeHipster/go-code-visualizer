package graph

import("fmt")

type PackageNodeRow struct {
	cells map[string]string
}

func CreatePackageNodeRow(fileName string, functions []string, types []string, variables []string) PackageNodeRow{
	
	pnr := PackageNodeRow{cells: make(map[string]string)}
	
	rows := 1
	if(len(functions) > rows){ rows = len(functions) }
	if(len(types) > rows){  rows = len(types) }
	if(len(variables) > rows){ rows = len(variables) }
	
	rowStr :=
`					%s<BR></BR>
`

	for i := 0 ; i < rows; i++{
		if(i == (rows-1) / 2 ){
			pnr.cells["filename"] += fmt.Sprintf(rowStr,fileName)	
		}else{
			pnr.cells["filename"] += fmt.Sprintf(rowStr,".")
		}
		
		if(i < len(functions)){
			pnr.cells["functions"] += fmt.Sprintf(rowStr, functions[i])
		}else{
			pnr.cells["functions"] += fmt.Sprintf(rowStr,".")
		}
		
		if(i < len(types)){
			pnr.cells["types"] += fmt.Sprintf(rowStr, types[i])
		}else{
			pnr.cells["types"] += fmt.Sprintf(rowStr,".")
		}
		
		if(i < len(variables)){
			pnr.cells["variables"] += fmt.Sprintf(rowStr, variables[i])
		}else{
			pnr.cells["variables"] += fmt.Sprintf(rowStr,".")
		}		
	}
	
	//remove last new line
	pnr.cells["filename"] = pnr.cells["filename"][:len(pnr.cells["filename"])-1] 
	pnr.cells["functions"] = pnr.cells["functions"][:len(pnr.cells["functions"])-1] 
	pnr.cells["types"] = pnr.cells["types"][:len(pnr.cells["types"])-1] 
	pnr.cells["variables"] = pnr.cells["variables"][:len(pnr.cells["variables"])-1] 
	
	return pnr
}

func (pnr PackageNodeRow) fileNameCell() string{	
	return pnr.cells["filename"]
}

func (pnr PackageNodeRow) functionsCell() string{
	return pnr.cells["functions"]
}

func (pnr PackageNodeRow) typesCell() string{
	return pnr.cells["types"]
}

func (pnr PackageNodeRow) variablesCell() string{
	return pnr.cells["variables"]
}
