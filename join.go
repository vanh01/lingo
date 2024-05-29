package lingo

import (
	"sync"

	"github.com/vanh01/lingo/definition"
)

// Join joins two sequences based on key selector functions and extracts pairs of values.
//
// In this method, comparer is returns whether left is equal to right or not.
//
// If comparer is empty or nil, we will use the default comparer.
// On the other hand, we just use the first comparer
func (e Enumerable[T]) Join(
	inner Enumerable[any],
	outerKeySelector definition.SingleSelector[T],
	innerKeySelector definition.SingleSelector[any],
	resultSelector definition.CombinationSelector[T, any],
	comparer ...definition.Comparer[any],
) Enumerable[any] {
	return Enumerable[any]{
		getIter: func() <-chan any {
			out := make(chan any)

			inners := inner.ToSlice()
			// use AsEnumerable to optimize performance, because inner can be a long chain of operators
			innerKeys := AsEnumerable(inners).Select(innerKeySelector).ToSlice()
			outers := e.ToSlice()
			// same inner
			outerKeys := AsEnumerable(outers).Select(outerKeySelector).ToSlice()

			go func() {
				defer close(out)
				for i := range outers {
					for j := range inners {
						if definition.IsEmptyOrNil(comparer) {
							if outerKeys[i] == innerKeys[j] {
								out <- resultSelector(outers[i], inners[j])
							}
						} else if comparer[0](outerKeys[i], innerKeys[j]) {
							out <- resultSelector(outers[i], inners[j])
						}
					}
				}
			}()

			return out
		},
	}
}

// ParallelEnumerable

// Join joins two parallel sequences based on key selector functions and extracts pairs of values.
//
// In this method, comparer is returns whether left is equal to right or not.
//
// If comparer is empty or nil, we will use the default comparer.
// On the other hand, we just use the first comparer
func (e ParallelEnumerable[T]) Join(
	inner ParallelEnumerable[any],
	outerKeySelector definition.SingleSelector[T],
	innerKeySelector definition.SingleSelector[any],
	resultSelector definition.CombinationSelector[T, any],
	comparer ...definition.Comparer[any],
) ParallelEnumerable[any] {
	return ParallelEnumerable[any]{
		wasSetUnordered: e.wasSetUnordered,
		ordered:         e.ordered,
		getIter: func() <-chan odata[any] {
			out := make(chan odata[any])

			inners := inner.ToSlice()
			// use AsEnumerable to optimize performance, because inner can be a long chain of operators
			innerKeys := AsEnumerable(inners).Select(innerKeySelector).ToSlice()
			outers := e.ToSlice()
			// same inner
			outerKeys := AsEnumerable(outers).Select(outerKeySelector).ToSlice()

			li := len(inners)

			go func() {
				defer close(out)
				var wg sync.WaitGroup
				for i := range outers {
					wg.Add(1)
					temp := outers[i]
					tempKey := outerKeys[i]
					tempI := i

					go func() {
						defer wg.Done()
						var wg1 sync.WaitGroup
						for j := range inners {
							wg1.Add(1)
							tempInner := inners[j]
							tempInnerKey := innerKeys[j]
							tempNo := tempI*li + j
							go func() {
								defer wg1.Done()
								if definition.IsEmptyOrNil(comparer) {
									if tempKey == tempInnerKey {
										out <- odata[any]{
											no:  tempNo,
											val: resultSelector(temp, tempInner),
										}
									}
								} else if comparer[0](tempKey, tempInnerKey) {
									out <- odata[any]{
										no:  tempNo,
										val: resultSelector(temp, tempInner),
									}
								}
							}()
						}
						wg1.Wait()
					}()
				}
				wg.Wait()
			}()

			return out
		},
	}
}
