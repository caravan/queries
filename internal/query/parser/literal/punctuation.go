package literal

import "github.com/caravan/kombi/parse"

// Punctuation parsers
var (
	Comma     = WS(parse.String(","))
	Semicolon = WS(parse.String(";"))
	Asterisk  = WS(parse.String("*"))
)
