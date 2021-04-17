package schema_test

import (
	"fmt"
	"testing"

	"github.com/caravan/queries/internal/query/schema"
	"github.com/caravan/queries/query"
	"github.com/caravan/streaming"
	"github.com/stretchr/testify/assert"
)

func TestStreamRegistry(t *testing.T) {
	as := assert.New(t)
	s := schema.New()
	as.NotNil(s)

	st := streaming.NewStream()
	err := s.RegisterStream("found", st)
	as.Nil(err)

	st2, ok := s.Stream("found")
	as.True(ok)
	as.Equal(st, st2)

	list := s.Streams()
	as.NotNil(list)
	as.Equal(1, len(list))
	as.Equal(query.SchemaName("found"), list[0])

	err = s.RegisterStream("found", streaming.NewStream())
	as.EqualError(err, fmt.Sprintf(schema.ErrItemAlreadyRegistered, "found"))
}

func TestStreamRegistryMissing(t *testing.T) {
	as := assert.New(t)

	s := schema.New()
	as.NotNil(s)

	as.Equal(0, len(s.Streams()))
	st, ok := s.Stream("missing")
	as.Nil(st)
	as.False(ok)
}
