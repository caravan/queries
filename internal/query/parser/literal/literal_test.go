package literal_test

import (
	"testing"

	"github.com/caravan/queries/internal/query/parser/literal"
	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	as := assert.New(t)

	s, f := literal.String.Parse(`'this is a string'`)
	as.NotNil(s)
	as.Nil(f)
	as.Equal("this is a string", s.Result)
}
