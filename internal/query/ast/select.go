package ast

import "github.com/caravan/queries/internal/query/lexer"

type (
	// SelectStatement represents a SQL SELECT Statement
	SelectStatement struct {
		lexer.Located
		ColumnSelectors
		SourceSelectors
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

	// ColumnSelectors is a set of ColumnSelector
	ColumnSelectors []*ColumnSelector

	// SourceSelectors is a set of SourceSelector
	SourceSelectors []*SourceSelector
)

// Statement marks SelectStatement as a Statement
func (*SelectStatement) Statement() {}
