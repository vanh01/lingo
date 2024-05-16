package lingo

import (
	"reflect"
	"sync"

	"github.com/vanh01/lingo/definition"
)

// Select projects values that are based on a transform function.
func (e Enumerable[T]) Select(selector definition.SingleSelector[T]) Enumerable[any] {
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
func (e Enumerable[T]) SelectMany(selector definition.SingleSelector[T]) Enumerable[any] {
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
func (e Enumerable[T]) Zip(second Enumerable[any], resultSelector ...definition.CombinationSelector[T, any]) Enumerable[any] {
	return Enumerable[any]{
		getIter: func() <-chan any {
			out := make(chan any)

			go func() {
				defer close(out)
				secondIter := second.getIter()
				for value := range e.getIter() {
					secondValue := <-secondIter
					if definition.IsEmptyOrNil(resultSelector) {
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

// ParallelEnumerable

// Select projects in parallel each element of a sequence into a new form.
func (p ParallelEnumerable[T]) Select(selector definition.SingleSelector[T]) ParallelEnumerable[any] {
	return ParallelEnumerable[any]{
		wasSetUnordered: p.wasSetUnordered,
		ordered:         p.ordered,
		getIter: func() <-chan odata[any] {
			out := make(chan odata[any])

			go func() {
				defer close(out)
				var wg sync.WaitGroup
				for value := range p.getIter() {
					wg.Add(1)
					temp := value
					go func() {
						var t any = selector(temp.val)
						out <- odata[any]{
							no:  temp.no,
							val: t,
						}
						wg.Done()
					}()
				}
				wg.Wait()
			}()

			return out
		},
	}
}

// SelectMany projects in parallel each element of a sequence to an []T and flattens the resulting sequences into one sequence.
func (p ParallelEnumerable[T]) SelectMany(selector definition.SingleSelector[T]) ParallelEnumerable[any] {
	return ParallelEnumerable[any]{
		wasSetUnordered: p.wasSetUnordered,
		ordered:         p.ordered,
		getIter: func() <-chan odata[any] {
			out := make(chan odata[any])

			go func() {
				defer close(out)
				var wg sync.WaitGroup

				for value := range p.getIter() {
					wg.Add(1)
					temp := value
					go func() {
						defer wg.Done()
						res := selector(temp.val)

						resValue := reflect.ValueOf(res)
						if resValue.Kind() == reflect.Slice {
							for i := 0; i < resValue.Len(); i++ {
								out <- odata[any]{
									no:  temp.no,
									val: resValue.Index(i).Interface(),
								}
							}
						}
					}()
				}
				wg.Wait()
			}()

			return out
		},
	}
}

// Zip merges in parallel two sequences by using the specified predicate function.
//
// If resultSelector is nil, the default result is a slice combined with each element
// On the other hand, we just use the first resultSelector
func (p ParallelEnumerable[T]) Zip(second ParallelEnumerable[any], resultSelector ...definition.CombinationSelector[T, any]) ParallelEnumerable[any] {
	return ParallelEnumerable[any]{
		wasSetUnordered: p.wasSetUnordered,
		ordered:         p.ordered,
		getIter: func() <-chan odata[any] {
			out := make(chan odata[any])

			go func() {
				defer close(out)
				var wg sync.WaitGroup
				secondIter := second.getIter()
				for value := range p.getIter() {
					wg.Add(1)
					secondValue := <-secondIter
					valueTemp := value
					secondTemp := secondValue

					go func() {
						defer wg.Done()
						if definition.IsEmptyOrNil(resultSelector) {
							out <- odata[any]{
								no:  valueTemp.no,
								val: []any{valueTemp.val, secondTemp.val},
							}
						} else {
							no := valueTemp.no
							valueTemp := resultSelector[0](valueTemp.val, secondTemp.val)
							out <- odata[any]{
								no:  no,
								val: valueTemp,
							}
						}
					}()
				}
				wg.Wait()
			}()

			return out
		},
	}
}
