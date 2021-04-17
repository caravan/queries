package schema_test

import (
	"fmt"
	"testing"

	"github.com/caravan/essentials"
	"github.com/caravan/queries/internal/query/schema"
	"github.com/caravan/queries/query"
	"github.com/stretchr/testify/assert"
)

func TestTopicRegistry(t *testing.T) {
	as := assert.New(t)
	s := schema.New()
	as.NotNil(s)

	st := essentials.NewTopic()
	err := s.RegisterTopic("found", st)
	as.Nil(err)

	st2, ok := s.Topic("found")
	as.True(ok)
	as.Equal(st, st2)

	list := s.Topics()
	as.NotNil(list)
	as.Equal(1, len(list))
	as.Equal(query.SchemaName("found"), list[0])

	err = s.RegisterTopic("found", essentials.NewTopic())
	as.EqualError(err, fmt.Sprintf(schema.ErrItemAlreadyRegistered, "found"))
}

func TestTopicRegistryMissing(t *testing.T) {
	as := assert.New(t)

	s := schema.New()
	as.NotNil(s)

	as.Equal(0, len(s.Topics()))
	st, ok := s.Topic("missing")
	as.Nil(st)
	as.False(ok)
}
