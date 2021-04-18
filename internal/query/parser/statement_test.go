package parser_test

import (
	"fmt"
	"testing"

	"github.com/caravan/queries/internal/query/lexer"
	"github.com/caravan/queries/internal/query/parser"
	"github.com/stretchr/testify/assert"
)

func testStatementsParse(t *testing.T, src string) {
	as := assert.New(t)
	seq := lexer.Lex(src)
	stmts, err := parser.Parse(seq).Statements()
	as.NotNil(stmts)
	as.Less(0, len(stmts))
	as.Nil(err)
}

func testStatementsError(t *testing.T, src, errStr string) {
	as := assert.New(t)
	seq := lexer.Lex(src)
	stmts, err := parser.Parse(seq).Statements()
	as.Nil(stmts)
	as.NotNil(err)
	as.Contains(err.Error(), errStr)
}

func testStatementParse(t *testing.T, src string) {
	as := assert.New(t)
	seq := lexer.Lex(src)
	res, err := parser.Parse(seq).Statement()
	as.NotNil(res)
	as.Nil(err)
}

func testStatementError(t *testing.T, src, errStr string) {
	as := assert.New(t)
	seq := lexer.Lex(src)
	res, err := parser.Parse(seq).Statement()
	as.Nil(res)
	as.NotNil(err)
	as.Contains(err.Error(), errStr)
}

func TestParsableStatements(t *testing.T) {
	testStatementsParse(t, "select column;")
	testStatementsParse(t, "select column alias;")
	testStatementsParse(t, "select column as alias;")
	testStatementsParse(t, "select column from table;")
	testStatementsParse(t, "select column alias from table;")
	testStatementsParse(t, "select column as col_alias from table;")
	testStatementsParse(t, "select column from table as tbl_alias;")
	testStatementsParse(t, "select col1, col2 from table1, table2")
	testStatementsParse(t, "select col from table where something;")
}

func TestParsableStatementsErrors(t *testing.T) {
	testStatementsError(t, "~;", parser.ErrExpectedStatement)
	testStatementsError(t, ";;~;", parser.ErrExpectedStatement)
	testStatementsError(t, "select;", parser.ErrExpectedExpression)
	testStatementsError(t, "select this from;", parser.ErrExpectedIdentifier)

	testStatementsError(t,
		"select 99 from this;",
		parser.ErrExpectedExpression,
	)

	testStatementsError(t,
		"select this from 99;",
		parser.ErrExpectedIdentifier,
	)

	testStatementsError(t,
		"select this from that as 99;",
		fmt.Sprintf(parser.ErrExpectedAlias, lexer.Integer),
	)

	testStatementsError(t,
		"select this from that where 99;",
		parser.ErrExpectedExpression,
	)
}

func TestSingleStatements(t *testing.T) {
	testStatementParse(t, "select this;")
	testStatementError(t, "", parser.ErrExpectedStatement)
}
