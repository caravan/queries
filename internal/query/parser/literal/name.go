package literal

import (
	"github.com/caravan/kombi/parse"
	"github.com/caravan/queries/internal/query/parser/ast"
)

// Name parses an atomic name
var Name = WS(parse.RegExp(`[a-zA-Z_][a-zA-Z0-9_@]*`)).Map(toName)

func nameMatcher(name ast.Name) func(ast.Name) bool {
	upper := toUpper(name)
	return func(n ast.Name) bool {
		return toUpper(n) == upper
	}
}

func toName(r parse.Result) parse.Result {
	return ast.Name(r.(string))
}
