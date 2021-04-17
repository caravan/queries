package ast_test

import (
	"testing"

	"github.com/caravan/queries/internal/query/ast"
	"github.com/stretchr/testify/assert"
)

func TestIdentifier(t *testing.T) {
	as := assert.New(t)
	i := &ast.Identifier{}
	e := ast.Expression(i)
	as.NotNil(e)
	e.Expression()
}
