package parser

import (
	"github.com/caravan/queries/internal/query/ast"
	"github.com/caravan/queries/internal/query/lexer"
)

// Error messages
const (
	ErrExpectedExpression = "expected expression"
)

// Expression parses a SQL expression
func (r *parser) Expression() (ast.Expression, error) {
	if res, err := r.expression(); res != nil || err != nil {
		return res, err
	}
	return nil, r.error(ErrExpectedExpression)
}

func (r *parser) expression() (ast.Expression, error) {
	res, err := r.identifier()
	if res != nil {
		return res, err
	}
	return nil, err
}

func (r *parser) Identifier() (*ast.Identifier, error) {
	res, err := r.identifier()
	if res != nil || err != nil {
		return res, err
	}
	return nil, r.error(ErrExpectedIdentifier)
}

func (r *parser) identifier() (*ast.Identifier, error) {
	r.pushState()
	t := r.nextToken()
	if !t.IsA(lexer.Identifier) {
		r.popState()
		return nil, nil
	}
	return &ast.Identifier{
		Located: t,
		Name:    t.Value().(string),
	}, nil
}
