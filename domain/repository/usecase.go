package repository

import "github.com/aimof/depfon/domain/entity"

// Analyzer analyzes FunctionTre
// args[0]: project root. example: "github.com/aimof/depfon"
// args[1]: entry function. example: "github.com/aimof/depfon/cmd/depfon.main"
// args[2]: options
// indludeVendor: include vendor directory to analyze. Default: exclude.
type Analyzer interface {
	Analyze(root, entry string, options uint) (*entity.FunctionTree, error)
}
