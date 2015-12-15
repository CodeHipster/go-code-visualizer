package graph

import(
	"fmt"
)

type PackageNode struct{
	Name string
	Path string
	Files []PackageNodeFile
}

type PackageNodeFile struct {
	FileName string
	Functions []string
	Types []string
	Variables []string
}

func (pnf PackageNodeFile) ToGraphString() map[string]string{
	
	packageColumnForFileMap := make(map[string]string)
	rows := 1
	if(len(pnf.Functions) > rows){ rows = len(pnf.Functions) }
	if(len(pnf.Types) > rows){  rows = len(pnf.Types) }
	if(len(pnf.Variables) > rows){ rows = len(pnf.Variables) }
	
	rowStr :=
`					%s<BR></BR>
`
	packageColumnForFileMap["functions"] = ""
	packageColumnForFileMap["types"] = ""
	packageColumnForFileMap["variables"] = ""
	
	for i := 0 ; i < rows; i++{
		packageColumnForFileMap["filename"] += fmt.Sprintf(rowStr,".")
		if(i < len(pnf.Functions)){
			packageColumnForFileMap["functions"] += fmt.Sprintf(rowStr, pnf.Functions[i])
		}else{
			packageColumnForFileMap["functions"] += fmt.Sprintf(rowStr,".")
		}
		
		if(i < len(pnf.Types)){
			packageColumnForFileMap["types"] += fmt.Sprintf(rowStr, pnf.Types[i])
		}else{
			packageColumnForFileMap["types"] += fmt.Sprintf(rowStr,".")
		}
		
		if(i < len(pnf.Variables)){
			packageColumnForFileMap["variables"] += fmt.Sprintf(rowStr, pnf.Variables[i])
		}else{
			packageColumnForFileMap["variables"] += fmt.Sprintf(rowStr,".")
		}		
	}
	return packageColumnForFileMap
}

//TODO: implement all the filenames, types, vars, funcs
func (pn PackageNode) ToGraphString() string {
	strNode := 
`	"%s" [
		label=<
		{	<B><FONT POINT-SIZE="12">%s : %s</FONT></B> |
%s			
		}>
	]`

	strBody :=
`			{
				{
					<B>FileName</B> |
%s
				}
				|
				{
					<B>Functions</B>|
%s					
				}
				|
				{
					<B>Types</B> |
%s										
				}
				|
				{
					<B>Variables</B>|
%s					
				}
			}`	
			
	strColumn :=
`%s
					|
`
	strFilenameColumn := ""
	strFuncColumn := ""
	strTypeColumn := ""
	strVarColumn := ""
	for _,graphFile := range pn.Files{
		mapFile := graphFile.ToGraphString()
		strFilenameColumn += fmt.Sprintf(strColumn, mapFile["filename"])
		strFuncColumn += fmt.Sprintf(strColumn, mapFile["functions"])
		strTypeColumn += fmt.Sprintf(strColumn, mapFile["types"])
		strVarColumn += fmt.Sprintf(strColumn, mapFile["variables"])
	}	
	
	nodeBody := fmt.Sprintf(strBody, strFilenameColumn, strFuncColumn, strTypeColumn, strVarColumn)
	
	return fmt.Sprintf(strNode, pn.Name, pn.Name, pn.Path, nodeBody)
}

