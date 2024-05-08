package lingo

// Join joins two sequences based on key selector functions and extracts pairs of values.
//
// In this method, comparer is returns whether left is equal to right or not.
//
// If comparer is empty or nil, we will use the default comparer.
// On the other hand, we just use the first comparer
func (e Enumerable[T]) Join(
	inner Enumerable[any],
	outerKeySelector SingleSelector[T],
	innerKeySelector SingleSelector[any],
	resultSelector CombinationSelector[T, any],
	comparer ...Comparer[any],
) Enumerable[any] {
	return Enumerable[any]{
		getIter: func() <-chan any {
			out := make(chan any)

			go func() {
				defer close(out)
				innerSlice := inner.ToSlice()
				for value := range e.getIter() {
					for _, i := range innerSlice {
						outerKey := outerKeySelector(value)
						innerKey := innerKeySelector(i)
						if isEmptyOrNil(comparer) {
							if outerKey == innerKey {
								out <- resultSelector(value, i)
							}
						} else if comparer[0](outerKey, innerKey) {
							out <- resultSelector(value, i)
						}
					}
				}
			}()

			return out
		},
	}
}
