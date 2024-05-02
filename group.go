package lingo

// GroupBy groups elements that share a common attribute.
//
// in this method, comparer is returns whether left is equal to right or not.
//
// elementSelector, comparer can be nil
func (e Enumerable[T]) GroupBy(
	keySelector SingleSelector[T],
	elementSelector SingleSelector[T],
	resultSelector GroupBySelector[any, any],
	getHash GetHashCode[any],
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
