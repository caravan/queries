package ast_test

import (
	"testing"

	"github.com/caravan/queries/internal/query/ast"
	"github.com/stretchr/testify/assert"
)

func TestSelectStatement(t *testing.T) {
	as := assert.New(t)
	sel := &ast.SelectStatement{}
	st := ast.Statement(sel)
	as.NotNil(st)
	st.Statement()

}
