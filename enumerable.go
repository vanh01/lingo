package lingo

import (
	"reflect"
	"unsafe"
)

type Enumerable[T any] struct {
	getIter func() <-chan T
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
					case isNumber(temp):
						out <- defaultConvertToNumber[T](value)
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
					case isNumber(temp):
						out <- defaultConvertToNumber[T](value)
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
