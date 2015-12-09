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
	fileName string
	functions []string
	types []string
	variables []string
}

//TODO: implement all the filenames, types, vars, funcs
func (pn PackageNode) ToGraphString() string {
	str := 
`	"%s" [
		label=<
		{	<B><FONT POINT-SIZE="12">%s : %s</FONT></B> |
			{
				{
					<B>FileName</B> |
				}
				|
				{
					<B>Functions</B>|
				}
				|
				{
					<B>Types</B> |					
				}
				|
				{
					<B>Variables</B>|
				}
			}
		}>
	]`
	return fmt.Sprintf(str, pn.Name, pn.Name, pn.Path)
}

