package schema

import (
	"sync"

	"github.com/caravan/essentials/topic"
	"github.com/caravan/queries/query"
)

type (
	topicRegistry struct {
		sync.RWMutex
		data topicData
	}

	topicData map[query.SchemaName]topic.Topic
)

func makeTopicRegistry() *topicRegistry {
	return &topicRegistry{
		data: topicData{},
	}
}

func (r *topicRegistry) List() []query.SchemaName {
	r.RLock()
	defer r.RUnlock()
	res := make([]query.SchemaName, 0, len(r.data))
	for k := range r.data {
		res = append(res, k)
	}
	return res
}

func (r *topicRegistry) Get(n query.SchemaName) (topic.Topic, bool) {
	r.RLock()
	defer r.RUnlock()
	res, ok := r.data[n]
	return res, ok
}

func (r *topicRegistry) Register(n query.SchemaName, t topic.Topic) error {
	r.Lock()
	defer r.Unlock()
	r.data[n] = t
	return nil
}
