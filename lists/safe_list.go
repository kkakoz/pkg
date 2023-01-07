package lists

import (
	"sync"
)

type safeList[T any] struct {
	list[T]

	lock sync.RWMutex
}

func (a *safeList[T]) Get(index int) (T, error) {
	a.lock.RLock()
	defer a.lock.RUnlock()
	return a.list.Get(index)
}

func (a *safeList[T]) AddLast(v T) {
	a.lock.Lock()
	defer a.lock.Unlock()
	a.list.AddLast(v)
}

func (a *safeList[T]) AddIndex(index int, v T) error {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.list.AddIndex(index, v)
}

func (a *safeList[T]) Delete(index int) error {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.list.Delete(index)
}

func (a *safeList[T]) Slice() []T {
	a.lock.RLock()
	defer a.lock.RUnlock()
	return a.list.Slice()
}

func (a *safeList[T]) Len() int {
	a.lock.RLock()
	defer a.lock.RUnlock()
	return a.list.Len()
}

func (a *safeList[T]) FilterBy(f func(T) bool) list[T] {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.list.FilterBy(f)
}

func (a *safeList[T]) Clone() list[T] {
	a.lock.RLock()
	defer a.lock.RUnlock()
	return a.list.Clone()
}

func (a *safeList[T]) String() string {
	a.lock.RLock()
	defer a.lock.RUnlock()
	return a.list.String()
}

func (a *safeList[T]) MarshalJSON() ([]byte, error) {
	a.lock.RLock()
	defer a.lock.RUnlock()
	return a.list.MarshalJSON()
}

func (a *safeList[T]) UnmarshalJSON(data []byte) error {
	a.lock.RLock()
	defer a.lock.RUnlock()
	return a.list.UnmarshalJSON(data)
}

var _ list[any] = (*safeList[any])(nil)

func NewSafeList[T any](l list[T]) list[T] {
	return &safeList[T]{
		list: l,
	}
}
