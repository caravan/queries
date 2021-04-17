package lexer

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"

	"github.com/caravan/queries/internal/query/reserved"
)

type (
	tokenizer func([]string) *Token

	matchEntry struct {
		pattern  *regexp.Regexp
		function tokenizer
	}

	matchEntries []matchEntry
)

// Error messages
const (
	ErrStringNotTerminated = "string has no closing quote"
	ErrUnexpectedCharacter = "unexpected character: %s"
	ErrExpectedFloat       = "value is not a float: %s"
	ErrExpectedInteger     = "value is not an integer: %s"

	errUnmatchedState = "unmatched lexing state"
)

var (
	escaped    = regexp.MustCompile(`''`)
	escapedMap = map[string]string{
		`''`: "'",
	}

	matchers = matchEntries{
		pattern(`$`, endState(EndOfFile)),

		pattern(`--[^\n]*([\n]|$)`, tokenState(EOLComment)),
		pattern(`/\*.*\*/`, tokenState(Comment)),
		pattern(`(\r\n|[\n\r])`, tokenState(NewLine)),
		pattern(`[\t\f ]+`, tokenState(Whitespace)),

		pattern(`(')(?P<s>([^']|'')*)('?)`, stringState),

		pattern(`[-]?\d*(\.\d+)?[eE][+-]?\d+`, floatState),
		pattern(`[-]?\d*\.\d+`, floatState),
		pattern(`[-]?\d+`, integerState),

		pattern(`,`, tokenState(Comma)),
		pattern(`;`, tokenState(Semicolon)),
		pattern(`\*`, tokenState(Asterisk)),

		pattern(`"([^"]|"")*"`, quotedIdentifierState),
		pattern("`([^`]|``)*`", quotedIdentifierState),
		pattern(`[a-zA-Z_][a-zA-Z0-9_@]*`, identifierState),

		pattern(`.`, errorState),
	}
)

// Lex creates a new lexer Sequence
func Lex(src string) Sequence {
	var r resolver
	var line, column int
	input := src

	r = func() (*Token, Sequence, bool) {
		t, rest := matchToken(input)
		if t.Type() != EndOfFile {
			t := t.WithLocation(line, column)

			if t.IsNewLine() {
				line++
				column = 0
			} else {
				column += len(input) - len(rest)
			}

			input = rest
			return t, newSequence(r), true
		}
		return t, nil, false
	}

	res := newSequence(r)
	return Filter(res, notWhitespace)
}

var notWhitespace = func(t *Token) bool {
	return !t.IsWhitespace()
}

func pattern(p string, s tokenizer) matchEntry {
	return matchEntry{
		pattern:  regexp.MustCompile("^" + p),
		function: s,
	}
}

func matchToken(src string) (*Token, string) {
	for _, s := range matchers {
		if sm := s.pattern.FindStringSubmatch(src); sm != nil {
			return s.function(sm), src[len(sm[0]):]
		}
	}
	// Programmer error
	panic(errors.New(errUnmatchedState))
}

func tokenState(t TokenType) tokenizer {
	return func(sm []string) *Token {
		return MakeToken(t, sm[0])
	}
}

func endState(t TokenType) tokenizer {
	return func(_ []string) *Token {
		return MakeToken(t, nil)
	}
}

func unescape(s string) string {
	r := escaped.ReplaceAllStringFunc(s, func(e string) string {
		return escapedMap[e]
	})
	return r
}

func stringState(sm []string) *Token {
	if len(sm[4]) == 0 {
		return MakeToken(Error, ErrStringNotTerminated)
	}
	s := unescape(sm[2])
	return MakeToken(String, s)
}

func floatState(sm []string) *Token {
	res, err := strconv.ParseFloat(sm[0], 64)
	if err != nil {
		errStr := fmt.Sprintf(ErrExpectedFloat, sm[0])
		return MakeToken(Error, errStr)
	}
	return MakeToken(Float, res)
}

func integerState(sm []string) *Token {
	res, err := strconv.ParseInt(sm[0], 10, 64)
	if err != nil {
		errStr := fmt.Sprintf(ErrExpectedInteger, sm[0])
		return MakeToken(Error, errStr)
	}
	return MakeToken(Integer, res)
}

func identifierState(sm []string) *Token {
	if id, ok := reserved.IsReserved(sm[0]); ok {
		return MakeToken(Reserved, id)
	}
	return MakeToken(Identifier, sm[0])
}

func quotedIdentifierState(sm []string) *Token {
	id := sm[0]
	return MakeToken(Identifier, id[1:len(id)-1])
}

func errorState(sm []string) *Token {
	errStr := fmt.Sprintf(ErrUnexpectedCharacter, sm[0])
	return MakeToken(Error, errStr)
}
