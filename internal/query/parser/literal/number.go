package literal

import (
	"strconv"

	"github.com/caravan/kombi/parse"
)

// Error messages
const (
	ErrExpectedFloat   = "value is not a float: %s"
	ErrExpectedInteger = "value is not an integer: %s"
)

// Float parses a floating point number
var Float = WS(parse.Any(
	parse.RegExp(`[-]?\d*(\.\d+)?[eE][+-]?\d+`),
	parse.RegExp(`[-]?\d*\.\d+`),
).Bind(func(r parse.Result) parse.Parser {
	s := r.(string)
	res, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return parse.Fail(ErrExpectedFloat, s)
	}
	return parse.Return(res)
}))

// Integer parses an integer
var Integer = WS(parse.
	RegExp(`[-]?\d+`).
	Bind(func(r parse.Result) parse.Parser {
		s := r.(string)
		res, err := strconv.ParseInt(s, 0, 64)
		if err != nil {
			return parse.Fail(ErrExpectedInteger, s)
		}
		return parse.Return(res)
	}))
