package ast

import "github.com/caravan/queries/internal/query/lexer"

type (
	// SelectStatement represents a SQL SELECT Statement
	SelectStatement struct {
		lexer.Located
		ColumnSelectors
		SourceSelectors
		*SelectCondition
	}

	// ColumnSelector represents a SQL column selector
	ColumnSelector struct {
		lexer.Located
		Expression
		Name string
	}

	// SourceSelector represents a SQL source selector (FROM)
	SourceSelector struct {
		lexer.Located
		Source string
		Name   string
	}

	SelectCondition struct {
		lexer.Located
		Expression // will be a BooleanExpression
	}

	// ColumnSelectors is a set of ColumnSelector
	ColumnSelectors []*ColumnSelector

	// SourceSelectors is a set of SourceSelector
	SourceSelectors []*SourceSelector
)

// Statement marks SelectStatement as a Statement
func (*SelectStatement) Statement() {}
