package schema

import (
	"sync"

	"github.com/caravan/queries/schema"
	"github.com/caravan/streaming/stream"
)

type (
	streamRegistry struct {
		sync.RWMutex
		data streamData
	}

	streamData map[schema.Name]stream.Stream
)

func makeStreamRegistry() *streamRegistry {
	return &streamRegistry{
		data: streamData{},
	}
}

func (r *streamRegistry) List() []schema.Name {
	r.RLock()
	defer r.RUnlock()
	res := make([]schema.Name, 0, len(r.data))
	for k := range r.data {
		res = append(res, k)
	}
	return res
}

func (r *streamRegistry) Get(n schema.Name) (stream.Stream, bool) {
	r.RLock()
	defer r.RUnlock()
	res, ok := r.data[n]
	return res, ok
}

func (r *streamRegistry) Register(n schema.Name, t stream.Stream) error {
	r.Lock()
	defer r.Unlock()
	r.data[n] = t
	return nil
}
