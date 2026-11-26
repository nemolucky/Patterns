package iterator

type Iterator[T any] interface {
	HasNext() bool
	Next() (T, bool)
}

type SliceIterator[T any] struct {
	slice []T
	index int
}

func NewSliceIterator[T any](slice []T) Iterator[T] {
	return &SliceIterator[T]{
		slice: slice,
		index: 0,
	}
}

func (si *SliceIterator[T]) HasNext() bool {
	return si.index < len(si.slice)
}

func (si *SliceIterator[T]) Next() (T, bool) {

	if !si.HasNext() {
		var zero T
		return zero, false
	}

	el := si.slice[si.index]
	si.index++

	return el, true
}
