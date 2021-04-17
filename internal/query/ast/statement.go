package ast

type (
	// Statement is any AST statement
	Statement interface {
		Node
		Statement()
	}

	// Statements represents a sequence of Statement nodes
	Statements []Statement
)
