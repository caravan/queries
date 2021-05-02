package schema

import (
	"sync"

	"github.com/caravan/essentials/topic"
	"github.com/caravan/queries/schema"
)

type (
	topicRegistry struct {
		sync.RWMutex
		data topicData
	}

	topicData map[schema.Name]topic.Topic
)

func makeTopicRegistry() *topicRegistry {
	return &topicRegistry{
		data: topicData{},
	}
}

func (r *topicRegistry) List() []schema.Name {
	r.RLock()
	defer r.RUnlock()
	res := make([]schema.Name, 0, len(r.data))
	for k := range r.data {
		res = append(res, k)
	}
	return res
}

func (r *topicRegistry) Get(n schema.Name) (topic.Topic, bool) {
	r.RLock()
	defer r.RUnlock()
	res, ok := r.data[n]
	return res, ok
}

func (r *topicRegistry) Register(n schema.Name, t topic.Topic) error {
	r.Lock()
	defer r.Unlock()
	r.data[n] = t
	return nil
}
