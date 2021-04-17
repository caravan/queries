package parser_test

import (
	"testing"

	"github.com/caravan/queries/internal/query/ast"
	"github.com/caravan/queries/internal/query/lexer"
	"github.com/caravan/queries/internal/query/parser"
	"github.com/stretchr/testify/assert"
)

func TestSelectTree(t *testing.T) {
	as := assert.New(t)

	lex := lexer.Lex(`select first_col as fc, second_col sc, third_col
					  from first_tbl as ft, second_tbl;`)

	p := parser.Parse(lex)
	s, err := p.Statement()
	as.NotNil(s)
	as.Nil(err)

	sel, ok := s.(*ast.SelectStatement)
	as.NotNil(sel)
	as.True(ok)

	as.Equal(3, len(sel.ColumnSelectors))
	as.Equal("fc", sel.ColumnSelectors[0].Name)
	as.Equal("sc", sel.ColumnSelectors[1].Name)
	as.Equal("third_col", sel.ColumnSelectors[2].Name)

	id, ok := sel.ColumnSelectors[0].Expression.(*ast.Identifier)
	as.True(ok)
	as.Equal("first_col", id.Name)

	id, ok = sel.ColumnSelectors[1].Expression.(*ast.Identifier)
	as.True(ok)
	as.Equal("second_col", id.Name)

	id, ok = sel.ColumnSelectors[2].Expression.(*ast.Identifier)
	as.True(ok)
	as.Equal("third_col", id.Name)

	as.Equal(2, len(sel.SourceSelectors))

	as.Equal("first_tbl", sel.SourceSelectors[0].Source)
	as.Equal("ft", sel.SourceSelectors[0].Name)

	as.Equal("second_tbl", sel.SourceSelectors[1].Source)
	as.Equal("second_tbl", sel.SourceSelectors[1].Name)
}
