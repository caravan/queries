package literal

import (
	"strings"

	"github.com/caravan/kombi/parse"
	"github.com/caravan/queries/internal/query/parser/ast"
)

// Reserved words
const (
	_as     = "AS"
	_from   = "FROM"
	_select = "SELECT"
	_where  = "WHERE"
)

// Error messages
const (
	ErrExpectedReservedWord = "expected reserved word: %s"
)

var reserved = makeMap([]ast.Name{
	_as, _from, _select, _where,
})

// Reserved word parsers
var (
	AS     = parseReserved(_as)
	FROM   = parseReserved(_from)
	SELECT = parseReserved(_select)
	WHERE  = parseReserved(_where)
)

// IsReserved returns whether or not the provided word is reserved
func IsReserved(name ast.Name) bool {
	res, ok := reserved[toUpper(name)]
	return ok && res
}

func makeMap(names []ast.Name) map[ast.Name]bool {
	res := make(map[ast.Name]bool, len(names))
	for _, name := range names {
		res[toUpper(name)] = true
	}
	return res
}

func toUpper(n ast.Name) ast.Name {
	return ast.Name(strings.ToUpper(string(n)))
}

func parseReserved(name ast.Name) parse.Parser {
	match := nameMatcher(name)
	return Name.Bind(func(r parse.Result) parse.Parser {
		s := r.(ast.Name)
		if IsReserved(s) && match(s) {
			return parse.Return(name)
		}
		return parse.Fail(ErrExpectedReservedWord, name)
	})
}
