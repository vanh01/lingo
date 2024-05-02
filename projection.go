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

type TFirst interface {
	Enumerable[any] | any
}

// Zip produces a sequence of tuples with elements from 2 specified sequences
//
// If resultSelector is nil, the default result is a slice combined with each element
func (e Enumerable[T]) Zip(first TFirst, resultSelector CombinationSelector[T, any]) Enumerable[any] {
	return Enumerable[any]{
		getIter: func() <-chan any {
			out := make(chan any)

			go func() {
				defer close(out)
				var firstSlice []any
				switch reflect.TypeOf(first).Kind() {
				case reflect.Slice:
					val := reflect.ValueOf(first)
					for j := 0; j < val.Len(); j++ {
						firstSlice = append(firstSlice, val.Index(j).Interface())
					}
				default:
					val := reflect.ValueOf(first).MethodByName("ToSlice").Call([]reflect.Value{})[0]
					for j := 0; j < val.Len(); j++ {
						firstSlice = append(firstSlice, val.Index(j).Interface())
					}
				}
				i := 0
				for value := range e.getIter() {
					if resultSelector == nil {
						out <- []any{value, firstSlice[i]}
					} else {
						out <- resultSelector(value, firstSlice[i])
					}
					i++
				}
			}()

			return out
		},
	}
}
