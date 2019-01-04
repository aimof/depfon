package analyzer

import "github.com/aimof/depfon/domain/entity"

// Analyzer analyze source code and make FunctionTree
type Analyzer struct {
	tree *entity.FunctionTree
}

// Analyze source code and return FunctionTree
func Analyze(root, entry string) (*entity.FunctionTree, error) {
	ana := newAnalyzer(root, entry)
	err := ana.analyze()
	if err != nil {
		return nil, err
	}
	return ana.tree, nil
}
