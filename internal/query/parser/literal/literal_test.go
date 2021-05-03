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

func TestFloat(t *testing.T) {
	as := assert.New(t)

	s, f := literal.Float(`32.98e+42`)
	as.NotNil(s)
	as.Nil(f)
	as.Equal(32.98e+42, s.Result)
}

func TestInteger(t *testing.T) {
	as := assert.New(t)

	s, f := literal.Integer(`4096`)
	as.NotNil(s)
	as.Nil(f)
	as.Equal(int64(4096), s.Result)
}
