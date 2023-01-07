package lists

import (
	"encoding/json"
	"fmt"
)

type arrayList[T any] struct {
	data []T
}

func (a *arrayList[T]) Get(index int) (T, error) {
	if index < 0 || index >= len(a.data) {
		return *new(T), fmt.Errorf("index out of range")
	}
	return a.data[index], nil
}

func (a *arrayList[T]) AddLast(v T) {
	a.data = append(a.data, v)
}

func (a *arrayList[T]) AddIndex(index int, v T) error {
	if index < 0 || index >= len(a.data) {
		return fmt.Errorf("index out of range")
	}
	left := a.data[:index]
	left = append(left, v)
	a.data = append(left, a.data[index:]...)
	return nil
}

func (a *arrayList[T]) Delete(index int) error {
	if index < 0 || index >= len(a.data) {
		return fmt.Errorf("index out of range")
	}
	left := a.data[:index]
	a.data = append(left, a.data[index+1:]...)
	return nil
}

func (a *arrayList[T]) Slice() []T {
	newData := make([]T, len(a.data))
	copy(newData, a.data)
	return newData
}

func (a *arrayList[T]) Len() int {
	return len(a.data)
}

func (a *arrayList[T]) FilterBy(f func(T) bool) list[T] {
	sub := make([]T, 0)
	for _, v := range a.data {
		if f(v) {
			sub = append(sub, v)
		}
	}
	return &arrayList[T]{
		data: sub,
	}
}

func (a *arrayList[T]) Clone() list[T] {
	return &arrayList[T]{
		data: a.Slice(),
	}
}

func (a *arrayList[T]) String() string {
	return fmt.Sprintf("%v", a.data)
}

func (a *arrayList[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.data)
}

func (a *arrayList[T]) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &a.data)
}

var _ list[any] = (*arrayList[any])(nil)

func NewArrayList[T any](opts ...option[T]) list[T] {
	opt := options[T]{}
	for _, o := range opts {
		opt = o(opt)
	}
	return &arrayList[T]{
		data: opt.data,
	}
}
