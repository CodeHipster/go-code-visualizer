package formatter

import(
	"github.com/thijsoostdam/go-code-visualizer/formatter/graph"
	_ "fmt"
	)

func GenerateDotGraph(goCode []ParsedCode) string{
	
	settings := graph.NodeSettings{FontSize:12,Shape:"record"}
	packageNodes := getPackageNodes(goCode)
	packageRelations := getPackageRelations(goCode)
	
	graph := graph.NewGraph(settings, packageNodes, packageRelations)	
	
	return graph.BuildGraphString()	
}

func getMapOfPackagePaths(goCode []ParsedCode) (packages map[string]string){
	
	packages = make(map[string]string)
	
	for _, code := range goCode{
		packages[code.PackagePath()] = code.PackageName()
	}	
	
	return packages
}

func getPackageNodes(goCode []ParsedCode) (packageNodes []graph.PackageNode){
	
	packagesMap := make(map[string]graph.PackageNode)
	
	//Map the files to their graph.PackageNode
	for _, parsedCode := range goCode{		
		packageNodeFile := graph.PackageNodeFile{
			FileName: parsedCode.FileName(),
			Functions: parsedCode.Functions(),
			Types: parsedCode.Types(),
			Variables: parsedCode.Variables()}
		
		node := packagesMap[parsedCode.PackageName()]
		node.Name = parsedCode.PackageName()
		node.Path = parsedCode.PackagePath()
		node.Files = append(node.Files, packageNodeFile)
		//maybe we should work with pointers?
		packagesMap[parsedCode.PackageName()] = node
	}
	
	//Take the graph.PackageNode out of the map.
	for _, packageNode := range packagesMap{
		packageNodes = append(packageNodes, packageNode)
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