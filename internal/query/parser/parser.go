package parser

import (
	"errors"
	"fmt"

	"github.com/caravan/queries/internal/query/ast"
	"github.com/caravan/queries/internal/query/lexer"
	"github.com/caravan/queries/internal/query/reserved"
)

type (
	// Parser is a stateful iteration interface for a Token stream that
	// is piloted by the FromScanner function and exposed as a Sequence
	Parser interface {
		Statements() (ast.Statements, error)
		Statement() (ast.Statement, error)
		Expression() (ast.Expression, error)
	}

	parserState struct {
		seq   lexer.Sequence
		token *lexer.Token
	}

	parser struct {
		parserState
		state []parserState
	}
)

// Error messages
const (
	ErrExpectedStatement = "expected statement"

	ErrExpectedWrapper = "%s, got %s"
)

// Parse creates a new SQL Parser using the provided lexer.Sequence
func Parse(seq lexer.Sequence) Parser {
	return &parser{
		parserState: parserState{
			seq: seq,
		},
		state: []parserState{},
	}
}

func (r *parser) nextToken() *lexer.Token {
	t, seq, ok := r.seq.Split()
	if !ok {
		t := lexer.MakeToken(lexer.EndOfFile, nil)
		r.token = t
		return t
	}
	r.seq = seq
	r.token = t
	return t
}

func (r *parser) pushState() {
	r.state = append(r.state, r.parserState)
}

func (r *parser) popState() {
	l := len(r.state)
	r.parserState = r.state[l-1]
	r.state = r.state[0 : l-1]
}

func (r *parser) expected(msg string) error {
	if r.token != nil {
		wrapped := fmt.Sprintf(ErrExpectedWrapper, msg, r.token.Type())
		return r.error(wrapped)
	}
	return r.error(msg)
}

func (r *parser) error(msg string) error {
	return r.wrapError(errors.New(msg))
}

func (r *parser) wrapError(err error) error {
	if t := r.token; t != nil {
		return t.WrapError(err)
	}
	return err
}

func (r *parser) commaSeparated() bool {
	r.pushState()
	if r.nextToken().IsA(lexer.Comma) {
		return true
	}
	r.popState()
	return false
}

func (r *parser) reserved(word string) (*lexer.Token, bool) {
	r.pushState()
	t := r.nextToken()
	if isReserved(t, word) {
		return t, true
	}
	r.popState()
	return nil, false
}

func isReserved(t *lexer.Token, w string) bool {
	if t == nil || !t.IsA(lexer.Reserved) {
		return false
	}
	tw, ok := reserved.IsReserved(t.Value().(string))
	return ok && tw == w
}
