package lexer_test

import (
	"testing"

	"github.com/caravan/queries/internal/query/lexer"
	"github.com/stretchr/testify/assert"
)

func TestTokenTypeString(t *testing.T) {
	as := assert.New(t)

	as.Equal("Identifier", lexer.Identifier.String())
	as.Equal("TokenType(99)", lexer.TokenType(99).String())
}
