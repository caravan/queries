package statement_test

import (
	"testing"

	"github.com/caravan/queries/internal/query/parser/ast"
	"github.com/caravan/queries/internal/query/parser/statement"
	"github.com/stretchr/testify/assert"
)

func TestSelectTree(t *testing.T) {
	as := assert.New(t)

	s, f := statement.Statement.Parse(
		`select first_col as fc, second_col sc, third_col
		 from first_tbl as ft, second_tbl
		 where blah;`,
	)

	as.NotNil(s)
	as.Nil(f)

	sel, ok := s.Result.(*ast.SelectStatement)
	as.NotNil(sel)
	as.True(ok)

	as.Equal(3, len(sel.ColumnSelectors))
	as.Equal(ast.Name("fc"), sel.ColumnSelectors[0].Name)
	as.Equal(ast.Name("sc"), sel.ColumnSelectors[1].Name)
	as.Equal(ast.Name("third_col"), sel.ColumnSelectors[2].Name)

	id, ok := sel.ColumnSelectors[0].Expression.(*ast.Identifier)
	as.True(ok)
	as.Equal(ast.Name("first_col"), id.Name)

	id, ok = sel.ColumnSelectors[1].Expression.(*ast.Identifier)
	as.True(ok)
	as.Equal(ast.Name("second_col"), id.Name)

	id, ok = sel.ColumnSelectors[2].Expression.(*ast.Identifier)
	as.True(ok)
	as.Equal(ast.Name("third_col"), id.Name)

	as.Equal(2, len(sel.SourceSelectors))

	as.Equal(ast.Name("first_tbl"), sel.SourceSelectors[0].Source)
	as.Equal(ast.Name("ft"), sel.SourceSelectors[0].Name)

	as.Equal(ast.Name("second_tbl"), sel.SourceSelectors[1].Source)
	as.Equal(ast.Name("second_tbl"), sel.SourceSelectors[1].Name)
}
