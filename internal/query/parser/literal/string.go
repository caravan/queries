package literal

import (
	"github.com/caravan/kombi/parse"
	"github.com/caravan/queries/internal/query/parser/literal/internal"
)

// String parses a string literal
var String = WS(internal.QuotedParser("'")).Map(toString)

func toString(r parse.Result) parse.Result {
	return string(r.(parse.Input))
}
