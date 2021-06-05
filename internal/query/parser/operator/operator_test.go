package operator_test

import (
	"testing"

	"github.com/caravan/queries/internal/query/parser/operator"
	"github.com/stretchr/testify/assert"
)

func TestRelationalString(t *testing.T) {
	as := assert.New(t)
	as.Equal("Greater Than", operator.GT.String())
	as.Equal("Relational(99)", operator.Relational(99).String())
}

func TestBinaryString(t *testing.T) {
	as := assert.New(t)
	as.Equal("Addition", operator.ADD.String())
	as.Equal("Binary(99)", operator.Binary(99).String())
}

func TestUnaryString(t *testing.T) {
	as := assert.New(t)
	as.Equal("Positive", operator.POS.String())
	as.Equal("Unary(99)", operator.Unary(99).String())
}
