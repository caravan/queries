package schema

import (
	"fmt"
	"sync"

	"github.com/caravan/essentials/topic"
	"github.com/caravan/queries/schema"
	"github.com/caravan/streaming/stream"
	"github.com/caravan/streaming/table"
)

type (
	// Schema is the internal implementation of a Schema
	Schema struct {
		sync.Mutex
		items   registered
		streams *streamRegistry
		tables  *tableRegistry
		topics  *topicRegistry
	}

	registered map[schema.Name]bool
)

// Error messages
const (
	ErrItemAlreadyRegistered = "item already registered in schema: %s"
)

// New return a new internal Schema instance
func New() schema.Schema {
	return &Schema{
		items:   registered{},
		streams: makeStreamRegistry(),
		tables:  makeTableRegistry(),
		topics:  makeTopicRegistry(),
	}
}

func (s *Schema) register(n schema.Name) error {
	s.Lock()
	defer s.Unlock()
	if res, ok := s.items[n]; ok && res {
		return fmt.Errorf(ErrItemAlreadyRegistered, n)
	}
	s.items[n] = true
	return nil
}

// Streams returns the list of Stream names in this Schema
func (s *Schema) Streams() []schema.Name {
	return s.streams.List()
}

// Stream returns a Stream from the Schema by name
func (s *Schema) Stream(n schema.Name) (stream.Stream, bool) {
	return s.streams.Get(n)
}

// RegisterStream registers a Stream with the Schema by name
func (s *Schema) RegisterStream(n schema.Name, st stream.Stream) error {
	if err := s.register(n); err != nil {
		return err
	}
	return s.streams.Register(n, st)
}

// Tables returns the list of Tables names in this Schema
func (s *Schema) Tables() []schema.Name {
	return s.tables.List()
}

// Table returns a Table from the Schema by name
func (s *Schema) Table(n schema.Name) (table.Table, bool) {
	return s.tables.Get(n)
}

// RegisterTable registers a Table with the Schema by name
func (s *Schema) RegisterTable(n schema.Name, t table.Table) error {
	if err := s.register(n); err != nil {
		return err
	}
	return s.tables.Register(n, t)
}

// Topics returns the list of Topic names in this Schema
func (s *Schema) Topics() []schema.Name {
	return s.topics.List()
}

// Topic returns a Topic from the Schema by name
func (s *Schema) Topic(n schema.Name) (topic.Topic, bool) {
	return s.topics.Get(n)
}

// RegisterTopic registers a Topic with the Schema by name
func (s *Schema) RegisterTopic(n schema.Name, t topic.Topic) error {
	if err := s.register(n); err != nil {
		return err
	}
	return s.topics.Register(n, t)
}
