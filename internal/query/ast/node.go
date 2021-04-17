package ast

import "github.com/caravan/queries/internal/query/lexer"

// Node is any AST node
type Node interface {
	lexer.Located
}
