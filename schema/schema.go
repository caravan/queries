package schema

import (
	"github.com/caravan/essentials/topic"
	"github.com/caravan/streaming/stream"
	"github.com/caravan/streaming/table"
)

type (
	// Name is the name of an item in the Schema
	Name string

	// Schema provides management and access to query resources
	Schema interface {
		Streams() []Name
		Stream(Name) (stream.Stream, bool)
		RegisterStream(Name, stream.Stream) error

		Tables() []Name
		Table(Name) (table.Table, bool)
		RegisterTable(Name, table.Table) error

		Topics() []Name
		Topic(Name) (topic.Topic, bool)
		RegisterTopic(Name, topic.Topic) error
	}
)
