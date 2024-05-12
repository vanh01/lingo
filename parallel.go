package lingo

// ParallelEnumerable provides a set of methods for querying objects that implement ParallelQuery[T].
// This is the parallel equivalent of Enumerable.
//
// # Note that ParallelEnumerable will be useful in some cases, you must really consider when using it.
//
// # In worse cases, not only will performance not be optimized, but costs will also increase.
type ParallelEnumerable[T any] struct {
	getIter func() <-chan T
}

// GetIter returns an unbuffered channel of T that iterates through a collection.
func (p ParallelEnumerable[T]) GetIter() <-chan T {
	return p.getIter()
}

// AsParallel creates a new ParallelEnumerable from an Enumerable, it will need more resources
// but performance will be improved
//
// # Note that ParallelEnumerable will be useful in some cases, you must really consider when using it.
//
// # In worse cases, not only will performance not be optimized, but costs will also increase.
func (e Enumerable[T]) AsParallel() ParallelEnumerable[T] {
	return ParallelEnumerable[T]{
		getIter: func() <-chan T {
			ch := make(chan T)

			go func() {
				defer close(ch)
				for value := range e.getIter() {
					ch <- value
				}
			}()

			return ch
		},
	}
}

// AsParallelEnumerable creates a new ParallelEnumerable, it will need more resources
// but performance will be improved
//
// # Note that ParallelEnumerable will be useful in some cases, you must really consider when using it.
//
// # In worse cases, not only will performance not be optimized, but costs will also increase.
func AsParallelEnumerable[T any](t []T) ParallelEnumerable[T] {
	return ParallelEnumerable[T]{
		getIter: func() <-chan T {
			ch := make(chan T)

			go func() {
				defer close(ch)
				for _, value := range t {
					ch <- value
				}
			}()

			return ch
		},
	}
}
