package graph

import(
	"fmt"
)

//TODO: look at which methods and types to export and which to keep private.

type Graph struct{
	settings []string
	nodeSetting NodeSettings
	packageNodes []PackageNode
	packageRelations []PackageRelation
}

func CreateGraph(settings []string, nodeSetting NodeSettings, packageNodes []PackageNode, packageRelations []PackageRelation) (graph Graph){
	graph.settings = settings
	graph.nodeSetting = nodeSetting
	graph.packageNodes = packageNodes
	graph.packageRelations = packageRelations
	
	return graph
}

func (g Graph)BuildGraphString() string {
	//print graph object and save the point where nodes can be added.
	graphString := 
`digraph GoProject {
%s		
%s
%s
%s
}`
	return fmt.Sprintf(graphString, g.settingsGraphString(), g.nodeSetting.ToGraphString(), g.packageNodesGraphString(), g.packageRelationsGraphString())
}

func (g Graph) settingsGraphString() string{
	var str string
	for _,setting := range g.settings{
		str += `	` + setting + "\n"
	}
	return str
}

func (g Graph) packageNodesGraphString() string{
	str := ""
	
	for _ ,node := range g.packageNodes{
		str += fmt.Sprintln(node.toGraphString()) 		
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





