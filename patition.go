package lingo

import (
	"sync"
)

// Skip skips elements up to a specified position in a sequence.
func (e Enumerable[T]) Skip(number int) Enumerable[T] {
	return Enumerable[T]{
		getIter: func() <-chan T {
			out := make(chan T)
			go func() {
				defer close(out)
				i := 1
				for value := range e.getIter() {
					if i > number {
						out <- value
					}
					i++
				}
			}()
			return out
		},
	}
}

// SkipWhile skips elements based on a predicate function until an element doesn't satisfy the condition.
func (e Enumerable[T]) SkipWhile(predicate Predicate[T]) Enumerable[T] {
	return Enumerable[T]{
		getIter: func() <-chan T {
			out := make(chan T)

			go func() {
				defer close(out)
				stopped := false
				for value := range e.getIter() {
					if !predicate(value) {
						stopped = true
					}
					if stopped {
						out <- value
					}
				}
			}()

			return out
		},
	}
}

// Take takes elements up to a specified position in a sequence.
func (e Enumerable[T]) Take(number int) Enumerable[T] {
	return Enumerable[T]{
		getIter: func() <-chan T {
			out := make(chan T)

			go func() {
				defer close(out)
				i := 0
				for value := range e.getIter() {
					if i < number {
						out <- value
					}
					i++
				}
			}()

			return out
		},
	}
}

// TakeWhile takes elements based on a predicate function until an element doesn't satisfy the condition.
func (e Enumerable[T]) TakeWhile(predicate Predicate[T]) Enumerable[T] {
	return Enumerable[T]{
		getIter: func() <-chan T {
			out := make(chan T)

			go func() {
				defer close(out)
				stopped := false
				for value := range e.getIter() {
					if !predicate(value) {
						stopped = true
					}
					if !stopped {
						out <- value
					}
				}
			}()

			return out
		},
	}
}

// Chunk splits the elements of a sequence into chunks of a specified maximum size.
func Chunk[T any](e Enumerable[T], size int) Enumerable[[]T] {
	return Enumerable[[]T]{
		getIter: func() <-chan []T {
			out := make(chan []T)

			go func() {
				defer close(out)
				chunk := []T{}
				i := 0
				for value := range e.getIter() {
					if i == size {
						out <- chunk
						chunk = []T{}
						i = 0
					}
					chunk = append(chunk, value)
					i++
				}
				out <- chunk
			}()

			return out
		},
	}
}

// ParallelEnumerable

// Skip skips a specified number of elements in a parallel sequence and then returns the remaining elements.
//
// If the source sequence is ordered, Skip skips first n elements.
// On the other hand, Skip skips any n elements.
func (p ParallelEnumerable[T]) Skip(number int) ParallelEnumerable[T] {
	return ParallelEnumerable[T]{
		wasSetUnordered: p.wasSetUnordered,
		ordered:         p.ordered,
		getIter: func() <-chan odata[T] {
			out := make(chan odata[T])

			go func() {
				defer close(out)
				if p.ordered {
					i := 1
					for value := range p.order().getIter() {
						if i > number {
							out <- value
						}
						i++
					}
					return
				}

				// in case the ParallelEnumerable is unordered
				i := 1
				var mu sync.Mutex
				var wg sync.WaitGroup
				for value := range p.getIter() {
					wg.Add(1)
					temp := value
					go func() {
						defer mu.Unlock()
						defer wg.Done()
						mu.Lock()
						if i > number {
							out <- temp
						}
						i++
					}()
				}
				wg.Wait()
			}()

			return out
		},
	}
}

// SkipWhile elements in a parallel sequence as long as a specified condition is true and then returns the remaining elements.
//
// If the source sequence is ordered, SkipWhile skips according to the ordered element.
// On the other hand, performs SkipWhile on the current arbitrary order.
func (p ParallelEnumerable[T]) SkipWhile(predicate Predicate[T]) ParallelEnumerable[T] {
	return ParallelEnumerable[T]{
		wasSetUnordered: p.wasSetUnordered,
		ordered:         p.ordered,
		getIter: func() <-chan odata[T] {
			out := make(chan odata[T])

			go func() {
				defer close(out)

				if p.ordered {
					stopped := false
					for value := range p.order().getIter() {
						if !stopped && !predicate(value.val) {
							stopped = true
						}
						if stopped {
							out <- value
						}
					}
					return
				}

				// in case the ParallelEnumerable is unordered
				stopped := false
				var mu sync.Mutex
				var wg sync.WaitGroup
				for value := range p.getIter() {
					temp := value
					wg.Add(1)
					go func() {
						defer mu.Unlock()
						defer wg.Done()
						mu.Lock()
						if !stopped && !predicate(temp.val) {
							stopped = true
						}
						if stopped {
							out <- temp
						}
					}()
				}
				wg.Wait()
			}()

			return out
		},
	}
}

// Take returns a specified number of contiguous elements from the start of a parallel sequence.
//
// If the source sequence is ordered, Take takes first n elements.
// On the other hand, Take takes any n elements.
func (p ParallelEnumerable[T]) Take(number int) ParallelEnumerable[T] {
	return ParallelEnumerable[T]{
		wasSetUnordered: p.wasSetUnordered,
		ordered:         p.ordered,
		getIter: func() <-chan odata[T] {
			out := make(chan odata[T])

			go func() {
				defer close(out)
				if p.ordered {
					temp := p.order()
					i := 0
					for value := range temp.getIter() {
						if i < number {
							out <- value
						}
						i++
					}
					return
				}

				// in case the ParallelEnumerable is unordered
				i := 0
				var mu sync.Mutex
				var wg sync.WaitGroup
				for value := range p.getIter() {
					wg.Add(1)
					temp := value
					go func() {
						defer mu.Unlock()
						defer wg.Done()
						mu.Lock()
						if i < number {
							out <- temp
						}
						i++
					}()
				}
				wg.Wait()
			}()

			return out
		},
	}
}

// TakeWhile takes elements in a parallel sequence based on a predicate function until an element doesn't satisfy the condition.
//
// If the source sequence is ordered, TakeWhile takes according to the ordered element.
// On the other hand, performs TakeWhile on the current arbitrary order.
func (p ParallelEnumerable[T]) TakeWhile(predicate Predicate[T]) ParallelEnumerable[T] {
	return ParallelEnumerable[T]{
		wasSetUnordered: p.wasSetUnordered,
		ordered:         p.ordered,
		getIter: func() <-chan odata[T] {
			out := make(chan odata[T])

			go func() {
				defer close(out)
				if p.ordered {
					stopped := false
					for value := range p.order().getIter() {
						if !stopped && !predicate(value.val) {
							stopped = true
						}
						if !stopped {
							out <- value
						}
					}
					return
				}

				// in case the ParallelEnumerable is unordered
				stopped := false
				var mu sync.Mutex
				var wg sync.WaitGroup
				for value := range p.getIter() {
					temp := value
					wg.Add(1)
					go func() {
						defer mu.Unlock()
						defer wg.Done()
						mu.Lock()
						if !stopped && !predicate(temp.val) {
							stopped = true
						}
						if !stopped {
							out <- temp
						}
					}()
				}
				wg.Wait()
			}()

			return out
		},
	}
}
