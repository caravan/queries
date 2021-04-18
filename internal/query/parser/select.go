package parser

import (
	"github.com/caravan/queries/internal/query/ast"
	"github.com/caravan/queries/internal/query/lexer"
	"github.com/caravan/queries/internal/query/reserved"
)

// Error message
const (
	ErrExpectedAlias      = "expected alias, got %s"
	ErrExpectedIdentifier = "expected identifier"
)

func (r *parser) selectStatement() (*ast.SelectStatement, error) {
	r.pushState()
	t := r.nextToken()

	if !isReserved(t, reserved.SELECT) {
		r.popState()
		return nil, nil
	}

	columns, err := r.columnSelectors(ast.ColumnSelectors{})
	if err != nil {
		return nil, err
	}

	sources, err := r.optionalSourceSelectors()
	if err != nil {
		return nil, err
	}

	condition, err := r.optionalCondition()
	if err != nil {
		return nil, err
	}

	return &ast.SelectStatement{
		Located:         t,
		ColumnSelectors: columns,
		SourceSelectors: sources,
		SelectCondition: condition,
	}, nil
}

func (r *parser) columnSelectors(
	res ast.ColumnSelectors,
) (ast.ColumnSelectors, error) {
	sel, err := r.columnSelector()
	if err != nil {
		return nil, err
	}
	return r.columnSelectorsRest(append(res, sel))
}

func (r *parser) columnSelectorsRest(
	res ast.ColumnSelectors,
) (ast.ColumnSelectors, error) {
	r.pushState()
	if !r.nextToken().IsA(lexer.Comma) {
		r.popState()
		return res, nil
	}
	return r.columnSelectors(res)
}

func (r *parser) columnSelector() (*ast.ColumnSelector, error) {
	exp, err := r.Expression()
	if exp == nil || err != nil {
		return nil, err
	}

	res := &ast.ColumnSelector{
		Located:    exp,
		Expression: exp,
		Name:       "",
	}

	if alias, err := r.selectorAlias(); alias != "" && err == nil {
		res.Name = alias
	} else if id, ok := exp.(*ast.Identifier); ok {
		res.Name = id.Name
	}

	return res, nil
}

func (r *parser) optionalSourceSelectors() (ast.SourceSelectors, error) {
	res := ast.SourceSelectors{}
	r.pushState()
	if !isReserved(r.nextToken(), reserved.FROM) {
		r.popState()
		return res, nil
	}
	return r.sourceSelectors(res)
}

func (r *parser) sourceSelectors(
	res ast.SourceSelectors,
) (ast.SourceSelectors, error) {
	sel, err := r.sourceSelector()
	if err != nil {
		return nil, err
	}
	return r.sourceSelectorsRest(append(res, sel))
}

func (r *parser) sourceSelectorsRest(
	res ast.SourceSelectors,
) (ast.SourceSelectors, error) {
	r.pushState()
	if !r.nextToken().IsA(lexer.Comma) {
		r.popState()
		return res, nil
	}
	return r.sourceSelectors(res)
}

func (r *parser) sourceSelector() (*ast.SourceSelector, error) {
	id, err := r.Identifier()
	if err != nil {
		return nil, err
	}

	res := &ast.SourceSelector{
		Located: id,
		Name:    id.Name,
		Source:  id.Name,
	}

	if alias, err := r.selectorAlias(); alias != "" && err == nil {
		res.Name = alias
	} else if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *parser) selectorAlias() (string, error) {
	r.pushState()
	t := r.nextToken()
	switch {
	case isReserved(t, reserved.AS):
		t = r.nextToken()
		if t.IsA(lexer.Identifier) {
			return t.Value().(string), nil
		}
		return "", r.errorf(ErrExpectedAlias, t.Type())
	case t.IsA(lexer.Identifier):
		return t.Value().(string), nil
	default:
		r.popState()
		return "", nil
	}
}

func (r *parser) optionalCondition() (*ast.SelectCondition, error) {
	r.pushState()
	if !isReserved(r.nextToken(), reserved.WHERE) {
		r.popState()
		return nil, nil
	}

	e, err := r.Expression()
	if err != nil {
		return nil, err
	}
	return &ast.SelectCondition{
		Located:    e,
		Expression: e,
	}, nil
}
