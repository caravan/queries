package lexer

import "fmt"

type (
	// TokenType is an opaque type for lexer tokens
	TokenType int

	// Token is a lexer value
	Token struct {
		Location
		typ   TokenType
		value interface{}
	}
)

//go:generate stringer -linecomment -output token_string.go -type TokenType
const (
	Error      TokenType = iota // Syntax Error
	Reserved                    // Reserved Word
	Identifier                  // Identifier
	String                      // String
	Float                       // Floating Point Number
	Integer                     // Integer
	Asterisk                    // Asterisk (*)
	Comma                       // Comma (,)
	Semicolon                   // Semicolon (;)
	Whitespace                  // Whitespace
	NewLine                     // New Line
	EOLComment                  // End of Line Comment
	Comment                     // Comment
	EndOfFile                   // End of File
)

// Error messages
const (
	ErrTokenWrapped = "%w (line %d, column %d)"
)

// MakeToken constructs a new scanner Token
func MakeToken(t TokenType, v interface{}) *Token {
	return &Token{
		typ:   t,
		value: v,
	}
}

// WithLocation returns a copy of the Token with location information
func (t *Token) WithLocation(line, column int) *Token {
	res := *t
	res.line = line
	res.column = column
	return &res
}

// Type returns the TokenType for this Token
func (t *Token) Type() TokenType {
	return t.typ
}

// IsA returns whether this Token is of a certain TokenType
func (t *Token) IsA(typ TokenType) bool {
	return t.typ == typ
}

// Value returns the scanned Value for this Token
func (t *Token) Value() interface{} {
	return t.value
}

// IsNewLine returns whether this Token represents a new line
func (t *Token) IsNewLine() bool {
	return t.typ == EOLComment || t.typ == NewLine
}

// IsWhitespace returns whether this Token represents whitespace
func (t *Token) IsWhitespace() bool {
	switch t.typ {
	case EOLComment, Comment, NewLine, Whitespace:
		return true
	default:
		return false
	}
}

// WrapError wraps an error with line and column information
func (t *Token) WrapError(e error) error {
	return fmt.Errorf(ErrTokenWrapped, e, t.line+1, t.column+1)
}
