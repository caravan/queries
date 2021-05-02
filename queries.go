package queries

import (
	"github.com/caravan/queries/schema"

	_schema "github.com/caravan/queries/internal/schema"
)

// NewSchema instantiates a new Schema
func NewSchema() schema.Schema {
	return _schema.New()
}
