package schema

import (
	"sync"

	"github.com/caravan/queries/schema"
	"github.com/caravan/streaming/table"
)

type (
	tableRegistry struct {
		sync.RWMutex
		data tableData
	}

	tableData map[schema.Name]table.Table
)

func makeTableRegistry() *tableRegistry {
	return &tableRegistry{
		data: tableData{},
	}
}

func (r *tableRegistry) List() []schema.Name {
	r.RLock()
	defer r.RUnlock()
	res := make([]schema.Name, 0, len(r.data))
	for k := range r.data {
		res = append(res, k)
	}
	return res
}

func (r *tableRegistry) Get(n schema.Name) (table.Table, bool) {
	r.RLock()
	defer r.RUnlock()
	res, ok := r.data[n]
	return res, ok
}

func (r *tableRegistry) Register(n schema.Name, t table.Table) error {
	r.Lock()
	defer r.Unlock()
	r.data[n] = t
	return nil
}
