package lexer_test

import (
	"fmt"
	"testing"

	"github.com/caravan/queries/internal/query/lexer"
	"github.com/caravan/queries/internal/query/reserved"
	"github.com/stretchr/testify/assert"
)

var eof = lexer.MakeToken(lexer.EndOfFile, nil)

func T(line int, col int, t lexer.TokenType, v interface{}) *lexer.Token {
	return lexer.MakeToken(t, v).WithLocation(line, col)
}

func assertToken(t *testing.T, like *lexer.Token, value *lexer.Token) {
	t.Helper()
	as := assert.New(t)
	as.Equal(like.Type(), value.Type())
	as.Equal(like.Value(), value.Value())
	as.Equal(like.Line(), value.Line())
	as.Equal(like.Column(), value.Column())
}

func assertTokenSequence(t *testing.T, s string, tokens []*lexer.Token) {
	t.Helper()
	as := assert.New(t)
	seq := lexer.Lex(s)
	var f *lexer.Token
	var r = seq
	var ok bool
	for _, l := range tokens {
		f, r, ok = r.Split()
		as.True(ok)
		assertToken(t, l, f)
	}
	f, r, ok = r.Split()
	assertToken(t, eof, f)
	as.False(ok)
	as.Nil(r)
}

func TestLexerBasics(t *testing.T) {
	assertTokenSequence(t, "select * from blah;", []*lexer.Token{
		T(0, 0, lexer.Reserved, reserved.SELECT),
		T(0, 7, lexer.Asterisk, "*"),
		T(0, 9, lexer.Reserved, reserved.FROM),
		T(0, 14, lexer.Identifier, "blah"),
		T(0, 18, lexer.Semicolon, ";"),
	})

	tokens := []*lexer.Token{
		T(0, 0, lexer.Reserved, reserved.SELECT),
		T(0, 7, lexer.Asterisk, "*"),
		T(0, 9, lexer.Reserved, reserved.FROM),
		T(0, 14, lexer.Identifier, "blah"),
		T(0, 20, lexer.Semicolon, ";"),
	}
	assertTokenSequence(t, "select * from `blah`;", tokens)
	assertTokenSequence(t, `select * from "blah";`, tokens)
}

func TestLexerNumbers(t *testing.T) {
	assertTokenSequence(t,
		"000 01 10 19.5 32.009 32.00 1e+10",
		[]*lexer.Token{
			T(0, 0, lexer.Integer, int64(0)),
			T(0, 4, lexer.Integer, int64(1)),
			T(0, 7, lexer.Integer, int64(10)),
			T(0, 10, lexer.Float, 19.5),
			T(0, 15, lexer.Float, 32.009),
			T(0, 22, lexer.Float, 32.0),
			T(0, 28, lexer.Float, 1e+10),
		})
}

func TestLexerStrings(t *testing.T) {
	assertTokenSequence(t, `
'this is the first string' -- eol string
--another eol string
'this is an '' escaped string'
'str1' /*comment*/ 'str2'
'unterminated string`,
		[]*lexer.Token{
			T(1, 0, lexer.String, "this is the first string"),
			T(3, 0, lexer.String, "this is an ' escaped string"),
			T(4, 0, lexer.String, "str1"),
			T(4, 19, lexer.String, "str2"),
			T(5, 0, lexer.Error, lexer.ErrStringNotTerminated),
		})
}

func TestLexerErrors(t *testing.T) {
	assertTokenSequence(t, `
100000000000000000000 ~
1000000000001e+90000000000040123123`,
		[]*lexer.Token{
			T(
				1, 0, lexer.Error,
				fmt.Sprintf(lexer.ErrExpectedInteger,
					"100000000000000000000"),
			),
			T(
				1, 22, lexer.Error,
				fmt.Sprintf(lexer.ErrUnexpectedCharacter, "~"),
			),
			T(
				2, 0, lexer.Error,
				fmt.Sprintf(lexer.ErrExpectedFloat,
					"1000000000001e+90000000000040123123"),
			),
		},
	)
}
