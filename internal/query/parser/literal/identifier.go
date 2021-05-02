package literal

import (
	"github.com/caravan/kombi/parse"
	"github.com/caravan/queries/internal/query/parser/ast"
	"github.com/caravan/queries/internal/query/parser/literal/internal"
)

// Error messages
const (
	ErrReservedWord = "reserved word: %s"
)

// Identifier parses an identifier Expression
var Identifier = parse.Any(
	WS(internal.QuotedParser(`"`)).Map(toName),
	WS(internal.QuotedParser("`")).Map(toName),
	Name.Bind(func(r parse.Result) parse.Parser {
		n := r.(ast.Name)
		if IsReserved(n) {
			return parse.Fail(ErrReservedWord, n)
		}
		return parse.Return(n)
	}),
).Map(func(r parse.Result) parse.Result {
	return &ast.Identifier{
		Name: r.(ast.Name),
	}
})
