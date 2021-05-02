package statement

import (
	"github.com/caravan/kombi/parse"
	"github.com/caravan/queries/internal/query/parser/ast"
	"github.com/caravan/queries/internal/query/parser/literal"
)

// Error messages
const (
	ErrStatementExpected = "statement expected"
)

// Statement parsers
var (
	Statements = Statement.ZeroOrMore().Combine(
		func(in ...parse.Result) parse.Result {
			out := make(ast.Statements, len(in))
			for i, s := range in {
				out[i] = s.(ast.Statement)
			}
			return out
		},
	).Bind(func(r parse.Result) parse.Parser {
		return EOF.Return(r).Or(parse.Fail(ErrStatementExpected))
	})

	Statement = parse.Any(
		SelectStatement,
		NonStatement,
	).Bind(func(r parse.Result) parse.Parser {
		return EOF.Or(literal.Semicolon).Return(r)
	})

	NonStatement = parse.Satisfy(
		func(i parse.Input) (int, error) {
			return len(i), nil
		},
	).Fail(ErrStatementExpected)

	EOF = literal.WS(parse.EOF)
)
