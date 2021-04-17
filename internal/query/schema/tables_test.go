package schema_test

import (
	"fmt"
	"testing"

	"github.com/caravan/queries/internal/query/schema"
	"github.com/caravan/queries/query"
	"github.com/caravan/streaming"
	"github.com/stretchr/testify/assert"
)

func TestTableRegistry(t *testing.T) {
	as := assert.New(t)
	s := schema.New()
	as.NotNil(s)

	st := streaming.NewTable(nil)
	err := s.RegisterTable("found", st)
	as.Nil(err)

	st2, ok := s.Table("found")
	as.True(ok)
	as.Equal(st, st2)

	list := s.Tables()
	as.NotNil(list)
	as.Equal(1, len(list))
	as.Equal(query.SchemaName("found"), list[0])

	err = s.RegisterTable("found", streaming.NewTable(nil))
	as.EqualError(err, fmt.Sprintf(schema.ErrItemAlreadyRegistered, "found"))
}

func TestTableRegistryMissing(t *testing.T) {
	as := assert.New(t)

	s := schema.New()
	as.NotNil(s)

	as.Equal(0, len(s.Tables()))
	st, ok := s.Table("missing")
	as.Nil(st)
	as.False(ok)
}
