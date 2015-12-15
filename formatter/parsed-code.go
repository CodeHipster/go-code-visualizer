package formatter


//Expects Imports to match PackagePath of other packages.
type ParsedCode interface{
	FileName() string
	//PackagePath() expects all paths to be lowercase.
	PackagePath() string
	PackageName() string
	//Imports() expects all paths to be lowercase.
	Imports() []string
	Types() []string
	Variables() []string
	Functions() []string
	ToString() string
}