package parser

import (
	"github.com/caravan/queries/internal/query/ast"
	"github.com/caravan/queries/internal/query/reserved"
)

// Error messages
const (
	ErrExpectedIdentifier = "expected identifier"
)

func (r *parser) selectStatement() (*ast.SelectStatement, error) {
	t, ok := r.reserved(reserved.SELECT)
	if !ok {
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
	if r.commaSeparated() {
		return r.columnSelectors(res)
	}
	return res, nil
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

	if alias, err := r.selectorAlias(); alias != nil && err == nil {
		res.Name = alias.Name
	} else if id, ok := exp.(*ast.Identifier); ok {
		res.Name = id.Name
	}

	return res, nil
}

func (r *parser) optionalSourceSelectors() (ast.SourceSelectors, error) {
	if _, ok := r.reserved(reserved.FROM); !ok {
		return nil, nil
	}
	return r.sourceSelectors(ast.SourceSelectors{})
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
	if r.commaSeparated() {
		return r.sourceSelectors(res)
	}
	return res, nil
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

	if alias, err := r.selectorAlias(); alias != nil && err == nil {
		res.Name = alias.Name
	} else if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *parser) selectorAlias() (*ast.Identifier, error) {
	if _, ok := r.reserved(reserved.AS); ok {
		return r.Identifier()
	}
	return r.identifier()
}

func (r *parser) optionalCondition() (*ast.SelectCondition, error) {
	if _, ok := r.reserved(reserved.WHERE); !ok {
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
