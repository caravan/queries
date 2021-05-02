package statement

import "github.com/caravan/kombi/parse"

// RequireIf returns a new Parser where the result is only parsed if the
// initial condition is met
func RequireIf(cond, result parse.Parser) parse.Parser {
	return parse.Or(
		cond.Return(condMet),
		parse.Return(condNotMet),
	).Bind(func(r parse.Result) parse.Parser {
		if r == condMet {
			return result
		}
		return parse.Return(nil)
	})
}
