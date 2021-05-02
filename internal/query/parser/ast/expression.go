package ast

import "github.com/caravan/queries/internal/query/parser/operator"

type (
	// Expression represents a SQL expression
	Expression interface {
		Node
		Expression()
	}

	// BinaryExpression represents an Expression with two operands
	BinaryExpression interface {
		Expression
		Left() Expression
		Right() Expression
	}

	// RelationalExpression is a comparative BinaryExpression
	RelationalExpression interface {
		BinaryExpression
		BooleanExpression
		Operator() operator.Relational
	}

	// BooleanExpression is an Expression that returns a Boolean result
	BooleanExpression interface {
		BooleanExpression()
	}

	// UnaryExpression is an Expression with only one operand
	UnaryExpression interface {
		Expression
		Operator() operator.Unary
	}

	// BooleanUnaryExpression is a Boolean Expression with one operand
	BooleanUnaryExpression interface {
		UnaryExpression
		BooleanExpression
	}

	// Identifier represents a SQL identifier
	Identifier struct {
		Name
	}
)

// Expression marks an Identifier as an Expression node
func (*Identifier) Expression() {}
