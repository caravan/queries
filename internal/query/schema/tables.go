package schema

import (
	"sync"

	"github.com/caravan/queries/query"
	"github.com/caravan/streaming/table"
)

type (
	tableRegistry struct {
		sync.RWMutex
		data tableData
	}

	tableData map[query.SchemaName]table.Table
)

func makeTableRegistry() *tableRegistry {
	return &tableRegistry{
		data: tableData{},
	}
}

func (r *tableRegistry) List() []query.SchemaName {
	r.RLock()
	defer r.RUnlock()
	res := make([]query.SchemaName, 0, len(r.data))
	for k := range r.data {
		res = append(res, k)
	}
	return res
}

func (r *tableRegistry) Get(n query.SchemaName) (table.Table, bool) {
	r.RLock()
	defer r.RUnlock()
	res, ok := r.data[n]
	return res, ok
}

func (r *tableRegistry) Register(n query.SchemaName, t table.Table) error {
	r.Lock()
	defer r.Unlock()
	r.data[n] = t
	return nil
}
