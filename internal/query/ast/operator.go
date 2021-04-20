package ast

//go:generate stringer -linecomment -output operator_string.go -type RelationalOperator,UnaryOperator

type (
	RelationalOperator int
	UnaryOperator      int
)

const (
	EQ  RelationalOperator = iota // Equal To
	NEQ                           // Not Equal To
	LT                            // Less Than
	LTE                           // Less Than Or Equal To
	GT                            // Greater Than
	GTE                           // Greater Than Or Equal To
	AND                           // And
	OR                            // Or
)

const (
	NEG UnaryOperator = iota // Negative
	POS                      // Positive
	NOT                      // Boolean Not
)
