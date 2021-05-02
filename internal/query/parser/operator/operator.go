package operator

import "github.com/caravan/kombi/parse"

//go:generate stringer -linecomment -output operator_string.go -type Relational,Binary,Unary

type (
	// Relational operator
	Relational int

	// Binary operator
	Binary int

	// Unary operator
	Unary int
)

// Relational operators
const (
	EQ  Relational = iota // Equal To
	NEQ                   // Not Equal To
	LT                    // Less Than
	LTE                   // Less Than Or Equal To
	GT                    // Greater Than
	GTE                   // Greater Than Or Equal To
	AND                   // And
	OR                    // Or
)

// Binary operators
const (
	ADD Binary = iota // Addition
	SUB               // Subtraction
	MUL               // Multiplication
	DIV               // Division
)

// Unary operators
const (
	NEG Unary = iota // Negative
	POS              // Positive
	NOT              // Boolean Not
)

// Operator parsers
var (
	RelationalOperator = parse.Any(
		parse.String("!=").Return(NEQ),
		parse.String("<>").Return(NEQ),
		parse.String("=").Return(EQ),
		parse.String("<=").Return(LTE),
		parse.String("<").Return(LT),
		parse.String(">=").Return(GTE),
		parse.String(">").Return(GT),
		parse.StrCaseCmp("AND").Return(AND),
		parse.StrCaseCmp("OR").Return(OR),
	)

	BinaryOperator = parse.Any(
		parse.String("+").Return(ADD),
		parse.String("-").Return(SUB),
		parse.String("*").Return(MUL),
		parse.String("/").Return(DIV),
	)

	UnaryOperator = parse.Any(
		parse.String("-").Return(NEG),
		parse.String("+").Return(POS),
		parse.StrCaseCmp("NOT").Return(NOT),
	)
)
