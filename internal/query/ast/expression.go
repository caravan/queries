package ast

import "github.com/caravan/queries/internal/query/lexer"

type (
	// Expression represents a SQL expression
	Expression interface {
		Node
		Expression()
	}

	BinaryExpression interface {
		Expression
		Left() Expression
		Right() Expression
	}

	RelationalExpression interface {
		BinaryExpression
		BooleanExpression
		Operator() RelationalOperator
	}

	BooleanExpression interface {
		BooleanExpression()
	}

	UnaryExpression interface {
		Expression
		Operator() UnaryOperator
	}

	BooleanUnaryExpression interface {
		UnaryExpression
		BooleanExpression
	}

	// Identifier represents a SQL identifier
	Identifier struct {
		lexer.Located
		Name string
	}
)

// Expression marks an Identifier as an Expression node
func (*Identifier) Expression() {}
