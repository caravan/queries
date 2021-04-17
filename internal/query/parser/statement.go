package parser

import (
	"github.com/caravan/queries/internal/query/ast"
	"github.com/caravan/queries/internal/query/lexer"
)

func (r *parser) Statement() (ast.Statement, error) {
	if res, err := r.statement(); res != nil || err != nil {
		return res, err
	}
	return nil, r.error(ErrExpectedStatement)
}

func (r *parser) Statements() (ast.Statements, error) {
	return r.statements(ast.Statements{})
}

func (r *parser) statement() (ast.Statement, error) {
	res, err := r.selectStatement()
	switch {
	case res != nil && err == nil:
		return res, nil
	case err != nil:
		return nil, err
	default:
		return nil, r.nonStatement()
	}
}

func (r *parser) nonStatement() error {
	r.pushState()
	for {
		switch r.nextToken().Type() {
		case lexer.Semicolon:
			continue
		case lexer.EndOfFile:
			return nil
		default:
			r.popState()
			return r.error(ErrExpectedStatement)
		}
	}
}

func (r *parser) statements(res ast.Statements) (ast.Statements, error) {
	s, err := r.statement()
	switch {
	case err != nil:
		return nil, err
	case s != nil:
		return r.statementsRest(append(res, s))
	default:
		return r.statementsRest(res)
	}
}

func (r *parser) statementsRest(res ast.Statements) (ast.Statements, error) {
	if r.nextToken().IsA(lexer.Semicolon) {
		return r.statements(res)
	}
	return res, nil
}
