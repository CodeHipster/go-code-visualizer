package graph

import("fmt")

type NodeSettings struct{
	FontSize int
	Shape string //FUTURE: make enumerable?
}

func(ns NodeSettings)ToGraphString() string{
	str :=	
`	node [
		fontsize = %d
		shape = "%s"
	]
`
	return fmt.Sprintf(str, ns.FontSize, ns.Shape)
}