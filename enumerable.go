package lingo

type Enumerable[T any] struct {
	iterator <-chan T
}

func lazyIterator[T any](slice []T) <-chan T {
	ch := make(chan T)

	go func() {
		defer close(ch)
		for _, value := range slice {
			ch <- value
		}
	}()

	return ch
}

// AsEnumerable creates a new Enumerable
func AsEnumerable[T any](t []T) Enumerable[T] {
	return Enumerable[T]{
		iterator: lazyIterator(t),
	}
}

// AsEnumerableTFromAny creates a new Enumerable of specific type from Enumerable of any
//
// This will be useful after using projection operations
func AsEnumerableTFromAny[T any](e Enumerable[any]) Enumerable[T] {
	out := make(chan T)

	go func() {
		defer close(out)
		for value := range e.iterator {
			out <- value.(T)
		}
	}()

	return Enumerable[T]{
		iterator: out,
	}
}

// AsEnumerableTFromSliceAny creates a new Enumerable of specific type from slice of any
//
// This will be useful after using projection operations
func AsEnumerableTFromSliceAny[T any](a []any) Enumerable[T] {
	out := make(chan T)

	go func() {
		defer close(out)
		for _, value := range a {
			out <- value.(T)
		}
	}()

	return Enumerable[T]{
		iterator: out,
	}
}
