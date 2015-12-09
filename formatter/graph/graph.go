package graph

import(
	"fmt"
)

//TODO: look at which methods and types to export and which to keep private.

type Graph struct{
	setting NodeSettings
	packageNodes []PackageNode
	packageRelations []PackageRelation
}

func NewGraph(setting NodeSettings, packageNodes []PackageNode, packageRelations []PackageRelation) (graph Graph){
	graph.setting = setting
	graph.packageNodes = packageNodes
	graph.packageRelations = packageRelations
	
	return graph
}

func (g Graph)BuildGraphString() string {
	//print graph object and save the point where nodes can be added.
	graphString := 
`digraph GoProject {
	rankdir=TB
	
%s
%s
%s
}`
	return fmt.Sprintf(graphString, g.setting.ToGraphString(), g.packageNodesGraphString(), g.packageRelationsGraphString())
}

func (g Graph) packageNodesGraphString() string{
	str := ""
	
	for _ ,node := range g.packageNodes{
		str += fmt.Sprintln(node.ToGraphString()) 		
	}
	
	return str
}

func (g Graph) packageRelationsGraphString() string{
	str := ""
	
	for _ ,node := range g.packageRelations{
		str += fmt.Sprintln(node.ToGraphString()) 		
	}
	
	return str
}





