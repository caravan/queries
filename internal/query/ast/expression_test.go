package ast_test

import (
	"testing"

	"github.com/caravan/queries/internal/query/ast"
)

func TestExpression(_ *testing.T) {
	(&ast.Identifier{}).Expression()
}
