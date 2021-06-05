package literal_test

import (
	"fmt"
	"testing"

	"github.com/caravan/kombi/parse"

	"github.com/caravan/queries/internal/query/parser/literal"
	"github.com/stretchr/testify/assert"
)

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

	tooBig := `4096000000000000000000000`
	s, f = literal.Integer(parse.Input(tooBig))
	as.Nil(s)
	as.NotNil(f)
	as.EqualError(f.Error, fmt.Sprintf(literal.ErrExpectedInteger, tooBig))
}
