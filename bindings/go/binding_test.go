package tree_sitter_tsql_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_tsql "github.com/krixgris/tree-sitter-tsql/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_tsql.Language())
	if language == nil {
		t.Errorf("Error loading Transact-SQL grammar")
	}
}
