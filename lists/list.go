package lists

type list[T any] interface {
	Get(index int) (T, error)
	AddLast(v T)
	AddIndex(index int, v T) error
	Delete(index int) error
	Len() int
	FilterBy(f func(T) bool) list[T]
	Clone() list[T]
	Slice() []T
	String() string
	MarshalJSON() ([]byte, error)
	UnmarshalJSON(data []byte) error
}

type options[T any] struct {
	data []T
}

type option[T any] func(options[T]) options[T]

func WithSlice[T any](slice []T) option[T] {
	return func(opts options[T]) options[T] {
		opts.data = slice
		return opts
	}
}
