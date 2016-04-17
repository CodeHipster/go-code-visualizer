package formatter

import(
	"github.com/codehipster/go-code-visualizer/formatter/graph"
	_ "fmt"
	)

func GenerateDotGraph(goCode []ParsedCode) string{
	
	settings := []string{"rankdir=TB"}
	nodeSettings := graph.NodeSettings{FontSize:12,Shape:"record"}
	packageNodes := getPackageNodes(goCode)
	packageRelations := getPackageRelations(goCode)
	
	graph := graph.CreateGraph(settings, nodeSettings, packageNodes, packageRelations)	
	
	return graph.BuildGraphString()	
}

func getPackageNodes(goCode []ParsedCode) (packageNodes []graph.PackageNode){
	
	packagesMap := make(map[string]*graph.PackageNode)
	
	//Map the files to their graph.PackageNode
	for _, parsedCode := range goCode{	
		
		//Get packageNode from map	
		nodePtr := packagesMap[parsedCode.PackageName()]
		if (nodePtr == nil){
			node := graph.CreatePackageNode(
				parsedCode.PackageName(),
				parsedCode.PackagePath(),
				nil)
			nodePtr = &node
		}
		
		//add row
		nodePtr.AddPackageNodeRow(graph.CreatePackageNodeRow(
			parsedCode.FileName(), 
			parsedCode.Functions(),
			parsedCode.Types(),
			parsedCode.Variables()))
		
		//put back in map
		packagesMap[parsedCode.PackageName()] = nodePtr
	}
	
	//Take the graph.PackageNode out of the map and put in an array.
	for _, packageNode := range packagesMap{
		packageNodes = append(packageNodes, *packageNode)
	}
	
	return packageNodes
}

func getPackageRelations(goCode []ParsedCode) (packageRelations []graph.PackageRelation){
	
	packagePaths := getMapOfPackagePaths(goCode)
	
	for _, file := range goCode{
		for _, dependencyPath := range file.Imports(){
			pkgname := packagePaths[dependencyPath]
			if(pkgname != ""){
				packageRelations = append(packageRelations, graph.PackageRelation{From: file.PackageName() , To: pkgname})
			}
		}
	}
	
	return packageRelations
}

func getMapOfPackagePaths(goCode []ParsedCode) (packages map[string]string){
	
	packages = make(map[string]string)
	
	for _, code := range goCode{
		packages[code.PackagePath()] = code.PackageName()
	}	
	
	return packages
}