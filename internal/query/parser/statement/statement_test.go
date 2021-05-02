package statement_test

import (
	"testing"

	"github.com/caravan/queries/internal/query/parser/ast"
	"github.com/caravan/queries/internal/query/parser/statement"
	"github.com/stretchr/testify/assert"
)

func testStatementsParse(t *testing.T, src string) {
	as := assert.New(t)
	s, f := statement.Statements.Parse(src)
	as.NotNil(s)
	as.Less(0, len(s.Result.(ast.Statements)))
	as.Nil(f)
}

func testStatementsError(t *testing.T, src, errStr string) {
	as := assert.New(t)
	s, f := statement.Statements.Parse(src)
	as.Nil(s)
	as.NotNil(f)
	as.Contains(f.Error.Error(), errStr)
}

func testStatementParse(t *testing.T, src string) {
	as := assert.New(t)
	s, f := statement.Statement.Parse(src)
	as.NotNil(s)
	as.Nil(f)
}

func testStatementError(t *testing.T, src, errStr string) {
	as := assert.New(t)
	s, f := statement.Statement.Parse(src)
	as.Nil(s)
	as.NotNil(f)
	as.Contains(f.Error.Error(), errStr)
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

var (
	statementError = statement.ErrStatementExpected
)

func TestParsableStatementsErrors(t *testing.T) {
	testStatementsError(t, "~;", statementError)
	testStatementsError(t, ";;~;", statementError)
	testStatementsError(t, "select;", statementError)
	testStatementsError(t, "select this from;", statementError)
	testStatementsError(t, "select 99 from this;", statementError)
	testStatementsError(t, "select this from 99;", statementError)
	testStatementsError(t, "select this from that as 99;", statementError)
	testStatementsError(t, "select this from that where 99;", statementError)
}

func TestSingleStatements(t *testing.T) {
	testStatementParse(t, "select this;")
	testStatementError(t, "", statementError)
}
