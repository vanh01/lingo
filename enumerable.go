package lingo

import (
	"reflect"
	"unsafe"

	"github.com/vanh01/lingo/definition"
)

type Enumerable[T any] struct {
	getIter func() <-chan T
}

// GetIter returns an unbuffered channel of T that iterates through a collection.
func (e Enumerable[T]) GetIter() <-chan T {
	return e.getIter()
}

// AsEnumerable creates a new Enumerable from ParallelEnumerable
func (p ParallelEnumerable[T]) AsEnumerable() Enumerable[T] {
	return Enumerable[T]{
		getIter: func() <-chan T {
			ch := make(chan T)

			temp := p
			if p.ordered {
				temp = temp.order()
			}
			go func() {
				defer close(ch)
				for value := range temp.getIter() {
					ch <- value.val
				}
			}()

			return ch
		},
	}
}

// AsEnumerable creates a new Enumerable
func AsEnumerable[T any](t []T) Enumerable[T] {
	return Enumerable[T]{
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

// AsEnumerableFromChannel creates a new Enumerable from a receive-only channel
func AsEnumerableFromChannel[T any](c <-chan T) Enumerable[T] {
	var t []T
	for value := range c {
		t = append(t, value)
	}

	return Enumerable[T]{
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

// AsEnumerableTFromAny creates a new Enumerable of specific type from Enumerable of any
//
// This will be useful after using projection operations
func AsEnumerableTFromAny[T any](e Enumerable[any]) Enumerable[T] {
	return Enumerable[T]{
		getIter: func() <-chan T {
			out := make(chan T)

			go func() {
				defer close(out)
				var temp T
				for value := range e.getIter() {
					switch {
					case definition.IsNumber(temp):
						out <- definition.DefaultConvertToNumber[T](value)
					case reflect.TypeOf(temp) == reflect.TypeOf(value):
						out <- value.(T)
					default:
						out <- *(*T)(unsafe.Pointer(&value))
					}
				}
			}()

			return out
		},
	}
}

// AsEnumerableAnyFromT creates a new Enumerable of any type from Enumerable of specific type
func AsEnumerableAnyFromT[T any](e Enumerable[T]) Enumerable[any] {
	return Enumerable[any]{
		getIter: func() <-chan any {
			out := make(chan any)

			go func() {
				defer close(out)
				for value := range e.getIter() {
					out <- value
				}
			}()

			return out
		},
	}
}

// AsEnumerableTFromSliceAny creates a new Enumerable of specific type from slice of any
//
// This will be useful after using projection operations
func AsEnumerableTFromSliceAny[T any](a []any) Enumerable[T] {

	return Enumerable[T]{
		getIter: func() <-chan T {
			out := make(chan T)

			go func() {
				defer close(out)
				var temp T
				for _, value := range a {
					switch {
					case definition.IsNumber(temp):
						out <- definition.DefaultConvertToNumber[T](value)
					case reflect.TypeOf(temp) == reflect.TypeOf(value):
						out <- value.(T)
					default:
						out <- *(*T)(unsafe.Pointer(&value))
					}
				}
			}()

			return out
		},
	}
}

// AsEnumerableAnyFromSliceT creates a new Enumerable of any type from slice of specific type
func AsEnumerableAnyFromSliceT[T any](a []T) Enumerable[any] {
	return Enumerable[any]{
		getIter: func() <-chan any {
			out := make(chan any)

			go func() {
				defer close(out)
				for _, value := range a {
					out <- value
				}
			}()

			return out
		},
	}
}

// Concat concatenates two sequences.
func (e Enumerable[T]) Concat(second Enumerable[T]) Enumerable[T] {
	return Enumerable[T]{
		getIter: func() <-chan T {
			out := make(chan T)

			go func() {
				defer close(out)
				for value := range e.getIter() {
					out <- value
				}
				for value := range second.getIter() {
					out <- value
				}
			}()

			return out
		},
	}
}

// Empty returns an empty Enumerable[T] that has the specified type argument.
func Empty[T any]() Enumerable[T] {
	return Enumerable[T]{
		getIter: func() <-chan T {
			out := make(chan T)

			go func() {
				defer close(out)
			}()

			return out
		},
	}
}

// Range generates a sequence of integral numbers within a specified range.
func Range(start, end int) Enumerable[int] {
	return Enumerable[int]{
		getIter: func() <-chan int {
			out := make(chan int)

			go func() {
				defer close(out)
				i := start
				for i <= end {
					out <- i
					i++
				}
			}()

			return out
		},
	}
}

// Repeat generates a sequence that contains one repeated value.
func Repeat[T any](element T, times int) Enumerable[T] {
	return Enumerable[T]{
		getIter: func() <-chan T {
			out := make(chan T)

			go func() {
				defer close(out)
				i := 0
				for i < times {
					out <- element
					i++
				}
			}()

			return out
		},
	}
}
