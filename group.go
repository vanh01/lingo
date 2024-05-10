package lingo

import "github.com/vanh01/lingo/definition"

// GroupBy groups elements that share a common attribute.
//
// In this method, getHash returns the key of the grouping.
//
// elementSelector, getHash can be nil
func (e Enumerable[T]) GroupBy(
	keySelector definition.SingleSelector[T],
	elementSelector definition.SingleSelector[T],
	resultSelector definition.GroupBySelector[any, any],
	getHash definition.GetHashCode[any],
) Enumerable[any] {
	return Enumerable[any]{
		getIter: func() <-chan any {
			out := make(chan any)

			go func() {
				defer close(out)
				res := map[any][]any{}
				for value := range e.getIter() {
					var element any
					key := keySelector(value)
					if elementSelector != nil {
						element = elementSelector(value)
					} else {
						element = value
					}
					if getHash != nil {
						key = getHash(key)
					}
					res[key] = append(res[key], element)
				}
				for k, v := range res {
					out <- resultSelector(k, v)
				}
			}()

			return out
		},
	}
}
