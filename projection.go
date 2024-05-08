package lingo

import (
	"reflect"
)

// Select projects values that are based on a transform function.
func (e Enumerable[T]) Select(selector SingleSelector[T]) Enumerable[any] {
	return Enumerable[any]{
		getIter: func() <-chan any {
			out := make(chan any)

			go func() {
				defer close(out)
				for value := range e.getIter() {
					out <- selector(value)
				}
			}()

			return out
		},
	}
}

// SelectMany projects sequences of values that are based on a transform function and then flattens them into one sequence.
func (e Enumerable[T]) SelectMany(selector SingleSelector[T]) Enumerable[any] {
	return Enumerable[any]{
		getIter: func() <-chan any {
			out := make(chan any)

			go func() {
				defer close(out)
				for value := range e.getIter() {
					res := selector(value)
					resValue := reflect.ValueOf(res)
					if resValue.Kind() == reflect.Slice {
						for i := 0; i < resValue.Len(); i++ {
							out <- resValue.Index(i).Interface()
						}
					}
				}
			}()

			return out
		},
	}
}

// Zip produces a sequence of tuples with elements from 2 specified sequences
//
// If resultSelector is nil, the default result is a slice combined with each element
// On the other hand, we just use the first resultSelector
func (e Enumerable[T]) Zip(second Enumerable[any], resultSelector ...CombinationSelector[T, any]) Enumerable[any] {
	return Enumerable[any]{
		getIter: func() <-chan any {
			out := make(chan any)

			go func() {
				defer close(out)
				secondIter := second.getIter()
				for value := range e.getIter() {
					secondValue := <-secondIter
					if isEmptyOrNil(resultSelector) {
						out <- []any{value, secondValue}
					} else {
						out <- resultSelector[0](value, secondValue)
					}
				}
			}()

			return out
		},
	}
}
