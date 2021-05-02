package internal

import (
	"fmt"
	"strings"

	"github.com/caravan/kombi/parse"
)

// Error messages
const (
	ErrValueNotTerminated = "value has no closing quote"
)

// QuotedParser returns a new Parser that is able to parse a quoted String
// or Identifier using the provided quotation mark
func QuotedParser(q string) parse.Parser {
	return parse.
		String(q).
		Then(parse.
			RegExp(escapedPattern(q)).
			Bind(func(r parse.Result) parse.Parser {
				return parse.Or(
					parse.String(q).Return(
						unescape(q, r.(string)),
					),
					parse.EOF.Fail(ErrValueNotTerminated),
				)
			}),
		)
}

func escapedPattern(q string) string {
	return fmt.Sprintf(`([^%s]|%s%s)*`, q, q, q)
}

func unescape(q string, s string) string {
	return strings.ReplaceAll(s, q+q, q)
}
