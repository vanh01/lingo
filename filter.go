package lingo

import "sync"

// Where selects values that are based on a predicate function.
func (e Enumerable[T]) Where(predicate Predicate[T]) Enumerable[T] {

	return Enumerable[T]{
		getIter: func() <-chan T {
			output := make(chan T)

			go func() {
				defer close(output)
				for value := range e.getIter() {
					if predicate(value) {
						output <- value
					}
				}
			}()

			return output
		},
	}
}

// ParallelEnumerable

// Where filters in parallel a sequence of values based on a predicate.
func (p ParallelEnumerable[T]) Where(predicate Predicate[T]) ParallelEnumerable[T] {
	return ParallelEnumerable[T]{
		wasSetUnordered: p.wasSetUnordered,
		ordered:         p.ordered,
		getIter: func() <-chan odata[T] {
			output := make(chan odata[T])

			go func() {
				defer close(output)
				var wg sync.WaitGroup
				for value := range p.getIter() {
					wg.Add(1)
					temp := value
					go func() {
						if predicate(temp.val) {
							output <- temp
						}
						wg.Done()
					}()
				}
				wg.Wait()
			}()

			return output
		},
	}
}
