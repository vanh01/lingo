package lingo

import (
	"reflect"
	"unsafe"
)

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
		var temp T
		for value := range e.iterator {
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

	return Enumerable[T]{
		iterator: out,
	}
}

// Concat concatenates two sequences.
func (e Enumerable[T]) Concat(second Enumerable[T]) Enumerable[T] {
	out := make(chan T)

	go func() {
		defer close(out)
		first := e.iterator
		secondI := second.iterator
		for value := range first {
			out <- value
		}
		for value := range secondI {
			out <- value
		}
	}()

	return Enumerable[T]{
		iterator: out,
	}
}

// Empty returns an empty Enumerable[T] that has the specified type argument.
func Empty[T any]() Enumerable[T] {
	out := make(chan T)

	go func() {
		defer close(out)
	}()

	return Enumerable[T]{
		iterator: out,
	}
}

// Range generates a sequence of integral numbers within a specified range.
func Range(start, end int) Enumerable[int] {
	out := make(chan int)

	go func() {
		defer close(out)
		i := start
		for i <= end {
			out <- i
			i++
		}
	}()

	return Enumerable[int]{
		iterator: out,
	}
}

// Repeat generates a sequence that contains one repeated value.
func Repeat[T any](element T, times int) Enumerable[T] {
	out := make(chan T)

	go func() {
		defer close(out)
		i := 0
		for i < times {
			out <- element
			i++
		}
	}()

	return Enumerable[T]{
		iterator: out,
	}
}
