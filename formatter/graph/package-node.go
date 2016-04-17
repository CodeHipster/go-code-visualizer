package graph

import(
	"fmt"
)

type PackageNode struct{
	name string
	path string
	rows []PackageNodeRow
}

func CreatePackageNode(fileName string, packagePath string, rows []PackageNodeRow) PackageNode{
	return PackageNode{
		name : fileName,
		path : packagePath,
		rows : rows}
}

func (pn *PackageNode) AddPackageNodeRow(row PackageNodeRow){
	pn.rows = append(pn.rows, row)
}

func (pn PackageNode) toGraphString() string {
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
	for _,row := range pn.rows{
		strFilenameColumn += fmt.Sprintf(strColumn, row.fileNameCell())
		strFuncColumn += fmt.Sprintf(strColumn, row.functionsCell())
		strTypeColumn += fmt.Sprintf(strColumn, row.typesCell())
		strVarColumn += fmt.Sprintf(strColumn, row.variablesCell())
	}	
	//remove the last piping symbol
	strFilenameColumn = strFilenameColumn[:len(strFilenameColumn) - 8]
	strFuncColumn = strFuncColumn[:len(strFuncColumn) - 8]
	strTypeColumn = strTypeColumn[:len(strTypeColumn) - 8]
	strVarColumn = strVarColumn[:len(strVarColumn) - 8]
	
	nodeBody := fmt.Sprintf(strBody, strFilenameColumn, strFuncColumn, strTypeColumn, strVarColumn)
	
	return fmt.Sprintf(strNode, pn.name, pn.name, pn.path, nodeBody)
}

