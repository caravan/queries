package ast

import "github.com/caravan/queries/internal/query/lexer"

type (
	// Expression represents a SQL expression
	Expression interface {
		Node
		Expression()
	}

	// Identifier represents a SQL identifier
	Identifier struct {
		lexer.Located
		Name string
	}
)

// Expression marks an Identifier as an Expression node
func (*Identifier) Expression() {}
