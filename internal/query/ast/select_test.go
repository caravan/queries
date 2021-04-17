package ast_test

import (
	"testing"

	"github.com/caravan/queries/internal/query/ast"
)

func TestSelectStatement(_ *testing.T) {
	(&ast.SelectStatement{}).Statement()
}
