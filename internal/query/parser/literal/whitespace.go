package literal

import "github.com/caravan/kombi/parse"

// Whitespace parsers
var (
	EOLComment = parse.RegExp(`--[^\n]*([\n]|$)`)
	Comment    = parse.RegExp(`/\*.*\*/`)
	NewLine    = parse.RegExp(`(\r\n|[\n\r])`)
	Whitespace = parse.RegExp(`[\t\f ]+`)

	AnyWS      = parse.Any(EOLComment, Comment, NewLine, Whitespace)
	OptionalWS = AnyWS.Optional()
)

// WS returns a new Parser that ignores leading and trailing whitespace
func WS(p parse.Parser) parse.Parser {
	return OptionalWS.Then(p).Bind(func(r parse.Result) parse.Parser {
		return OptionalWS.Return(r)
	})
}
