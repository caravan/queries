package query

import (
	"github.com/caravan/essentials/topic"
	"github.com/caravan/streaming/stream"
	"github.com/caravan/streaming/table"
)

type (
	// SchemaName is the name of an item in the Schema
	SchemaName string

	// Schema provides management and access to query resources
	Schema interface {
		Streams() []SchemaName
		Stream(SchemaName) (stream.Stream, bool)
		RegisterStream(SchemaName, stream.Stream) error

		Tables() []SchemaName
		Table(SchemaName) (table.Table, bool)
		RegisterTable(SchemaName, table.Table) error

		Topics() []SchemaName
		Topic(SchemaName) (topic.Topic, bool)
		RegisterTopic(SchemaName, topic.Topic) error
	}
)
