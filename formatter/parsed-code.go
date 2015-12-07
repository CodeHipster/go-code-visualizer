package formatter


//Expects Imports to match PackagePath of other packages.
type ParsedCode interface{
	PackagePath() string
	PackageName() string
	Imports() []string
	Types() []string
	Variables() []string
	Functions() []string
	ToString() string
}