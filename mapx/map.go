package mapx

import (
	"golang.org/x/exp/constraints"
	"sync"
)

type SafeMap[K constraints.Ordered, V any] struct {
	m    map[K]V
	lock sync.RWMutex
}

func NewSafeMap[K constraints.Ordered, V any]() *SafeMap[K, V] {
	return &SafeMap[K, V]{m: map[K]V{}}
}

func (m *SafeMap[K, V]) Add(k K, v V) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.m[k] = v
}

func (m *SafeMap[K, V]) Get(k K) (V, bool) {
	m.lock.RUnlock()
	defer m.lock.RUnlock()
	v, ok := m.m[k]
	return v, ok
}

func (m *SafeMap[K, V]) Delete(k K) {
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.m, k)
}

func (m *SafeMap[K, V]) Len() int {
	m.lock.RUnlock()
	defer m.lock.RUnlock()
	return len(m.m)
}

func (m *SafeMap[K, V]) Values() []V {
	m.lock.RUnlock()
	defer m.lock.RUnlock()
	res := make([]V, 0, len(m.m))
	for _, v := range m.m {
		res = append(res, v)
	}
	return res
}
