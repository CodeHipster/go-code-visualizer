package formatter

import(
	"github.com/thijsoostdam/go-code-visualizer/formatter/graph"
	"fmt"
	)

func GenerateDotGraph(goCode []ParsedCode) string{
	
	settings := graph.NodeSettings{FontSize:12,Shape:"record"}
	packageNodes := getPackageNodes(goCode)
	packageRelations := getPackageRelations(goCode)
	
	graph := graph.NewGraph(settings, packageNodes, packageRelations)	
	
	return graph.BuildGraphString()	
}

func getMapOfPackageNames(goCode []ParsedCode) (packages map[string]string){
	
	packages = make(map[string]string)
	
	for _, code := range goCode{
		packages[code.PackageName()] = code.PackagePath()
	}	
	
	return packages
}

func getMapOfPackagePaths(goCode []ParsedCode) (packages map[string]string){
	
	packages = make(map[string]string)
	
	for _, code := range goCode{
		packages[code.PackagePath()] = code.PackageName()
	}	
	
	return packages
}

func getPackageNodes(goCode []ParsedCode) (packageNodes []graph.PackageNode){
	
	packageNames := getMapOfPackageNames(goCode)
	
	for name, path := range packageNames {
		packageNodes = append(packageNodes, graph.PackageNode{Name:name, Path:path, Files:nil})
	}
	
	return packageNodes
}

func getPackageRelations(goCode []ParsedCode) (packageRelations []graph.PackageRelation){
	
	packagePaths := getMapOfPackagePaths(goCode)
	
	fmt.Printf("package paths:\n%v\n",packagePaths)
	
	for _, file := range goCode{
		fmt.Printf("file name: %s\n", file.PackageName())
		for _, dependencyPath := range file.Imports(){
			fmt.Printf("	dependency: %s\n",dependencyPath)
			pkgname := packagePaths[dependencyPath]
			if(pkgname != ""){
				packageRelations = append(packageRelations, graph.PackageRelation{From: file.PackageName() , To: pkgname})
			}
		}
	}
	
	return packageRelations
}