package lingo

// Join joins two sequences based on key selector functions and extracts pairs of values.
//
// in this method, comparer is returns whether left is equal to right or not.
func (e Enumerable[T]) Join(
	inner Enumerable[any],
	outerKeySelector SingleSelector[T],
	innerKeySelector SingleSelector[any],
	resultSelector CombinationSelector[T, any],
	comparer Comparer[any],
) Enumerable[any] {
	out := make(chan any)

	go func() {
		defer close(out)
		innerSlice := inner.ToSlice()
		for value := range e.iterator {
			for _, i := range innerSlice {
				outerKey := outerKeySelector(value)
				innerKey := innerKeySelector(i)
				if comparer == nil {
					if outerKey == innerKey {
						out <- resultSelector(value, i)
					}
				} else if comparer(outerKey, innerKey) {
					out <- resultSelector(value, i)
				}
			}
		}
	}()

	return Enumerable[any]{
		iterator: out,
	}
}
