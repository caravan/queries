package schema

import (
	"sync"

	"github.com/caravan/queries/query"
	"github.com/caravan/streaming/stream"
)

type (
	streamRegistry struct {
		sync.RWMutex
		data streamData
	}

	streamData map[query.SchemaName]stream.Stream
)

func makeStreamRegistry() *streamRegistry {
	return &streamRegistry{
		data: streamData{},
	}
}

func (r *streamRegistry) List() []query.SchemaName {
	r.RLock()
	defer r.RUnlock()
	res := make([]query.SchemaName, 0, len(r.data))
	for k := range r.data {
		res = append(res, k)
	}
	return res
}

func (r *streamRegistry) Get(n query.SchemaName) (stream.Stream, bool) {
	r.RLock()
	defer r.RUnlock()
	res, ok := r.data[n]
	return res, ok
}

func (r *streamRegistry) Register(n query.SchemaName, t stream.Stream) error {
	r.Lock()
	defer r.Unlock()
	r.data[n] = t
	return nil
}
