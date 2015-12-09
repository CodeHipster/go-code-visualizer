package graph

import("fmt")

type PackageRelation struct{
	From string
	To string
}

func (pr PackageRelation) ToGraphString() string {
	str := 
`	"%s" -> "%s"`
	return fmt.Sprintf(str, pr.From, pr.To)
}