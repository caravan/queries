package expression

import (
	"github.com/caravan/kombi/parse"
	"github.com/caravan/queries/internal/query/parser/literal"
)

// Error messages
const (
	ErrExpectedExpression = "expected expression"
)

// Expression parses an expression to be used in column projection
// or filtering conditions
var Expression = parse.Any(
	literal.Identifier,
)
